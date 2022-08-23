package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"strings"
	"trojan-panel-core/app/trojango"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/service"
	"trojan-panel-core/util"
)

type TrojanGoProcess struct {
	process
}

func NewTrojanGoInstance() *TrojanGoProcess {
	return &TrojanGoProcess{process{mutex: &mutex, binaryType: 2, cmdMap: &cmdMap}}
}

func (t *TrojanGoProcess) StartTrojanGo(apiPort uint) error {
	defer t.mutex.Unlock()
	if t.mutex.TryLock() {
		if t.IsRunning(apiPort) {
			return nil
		}
		binaryFilePath, err := util.GetBinaryFile(2)
		if err != nil {
			return err
		}
		configFilePath, err := util.GetConfigFile(2, apiPort)
		if err != nil {
			return err
		}
		cmd := exec.Command(binaryFilePath, "-config", configFilePath)
		t.cmdMap.Store(apiPort, cmd)
		runtime.SetFinalizer(t, t.Stop(apiPort))
		if err := cmd.Start(); err != nil {
			logrus.Errorf("start trojan-go error err: %v\n", err)
			return errors.New(constant.TrojanGoStartError)
		}
		go t.handlerUserUploadAndDownload(apiPort)
		return nil
	}
	logrus.Errorf("start trojan-go error err: lock not acquired\n")
	return errors.New(constant.TrojanGoStartError)
}

// 更新数据库中用户的下载和上传流量
func (t *TrojanGoProcess) handlerUserUploadAndDownload(apiPort uint) {
	api := trojango.NewTrojanGoApi(apiPort)
	for {
		if !t.IsRunning(apiPort) {
			logrus.Errorf("数据库同步至Trojan Go apiPort: %d trojan go not running\n", apiPort)
			break
		}
		users, err := api.ListUsers()
		if err != nil {
			continue
		}
		for _, user := range users {
			downloadTraffic := int(user.GetTrafficTotal().GetDownloadTraffic())
			uploadTraffic := int(user.GetTrafficTotal().GetDownloadTraffic())
			password := user.GetUser().GetPassword()
			passwordSplit := strings.Split(password, "&")
			if len(passwordSplit) != 2 || len(passwordSplit[0]) == 0 {
				continue
			}
			if err := service.UpdateAccountFlowByUsername(passwordSplit[0], downloadTraffic,
				uploadTraffic); err != nil {
				logrus.Errorf("Trojan Go同步至数据库 apiPort: %d 更新用户失败 err: %v", apiPort, err)
				continue
			}
			if err := api.DeleteUser(password); err != nil {
				logrus.Errorf("Trojan Go同步至数据库 apiPort: %d 删除用户失败 err: %v", apiPort, err)
				continue
			}
		}
	}
}
