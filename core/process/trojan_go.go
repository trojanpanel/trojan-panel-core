package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"sync"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/pkg/trojango"
	"trojan-panel-core/service"
	"trojan-panel-core/util"
)

type TrojanGoProcess struct {
	process
}

func NewTrojanGoProcess(apiPort string) (*TrojanGoProcess, error) {
	var mutex sync.Mutex
	defer mutex.Unlock()
	if mutex.TryLock() {
		t := &TrojanGoProcess{}
		binaryFilePath, err := util.GetBinaryFile("trojan-go")
		if err != nil {
			return nil, err
		}
		configFilePath, err := util.GetConfigFile(apiPort, "trojan-go")
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(binaryFilePath, "-config", configFilePath)
		t.cmdMap.Store(apiPort, cmd)
		runtime.SetFinalizer(t, t.Stop(apiPort))
		return t, nil
	}
	logrus.Errorf("new trojan-go process errror err: lock not acquired\n")
	return nil, errors.New(constant.NewTrojanGoProcessError)
}

func (t *TrojanGoProcess) StartTrojanGo(apiPort string) error {
	defer t.mutex.Unlock()
	if t.mutex.TryLock() {
		if t.IsRunning(apiPort) {
			return nil
		}
		cmd, ok := t.cmdMap.Load(apiPort)
		if ok {
			if err := cmd.(*exec.Cmd).Start(); err != nil {
				logrus.Errorf("start trojan-go error err: %v\n", err)
				return errors.New(constant.TrojanGoStartError)
			}
			t.handlerUsers(apiPort)
			go t.handlerUserUploadAndDownload(apiPort)
			go t.removeUsers(apiPort)
			return nil
		}
		logrus.Errorf("start trojan-go error err: process not found\n")
		return errors.New(constant.TrojanGoStartError)
	}
	logrus.Errorf("start trojan-go error err: lock not acquired\n")
	return errors.New(constant.TrojanGoStartError)
}

func (t *TrojanGoProcess) handlerUsers(apiPort string) {
	addApiUserVos, err := service.SelectUsersPassword(true)
	if err != nil {
		logrus.Errorf("查询api添加用户失败 err: %v\n", err)
		return
	}
	api := trojango.NewTrojanGoApi(apiPort)
	for _, apiUserVo := range addApiUserVos {
		if err := api.AddUser(dto.TrojanGoAddUserDto{
			Password:           apiUserVo.Password,
			UploadTraffic:      apiUserVo.Upload,
			DownloadTraffic:    apiUserVo.Download,
			IpLimit:            0,
			DownloadSpeedLimit: 0,
			UploadSpeedLimit:   0,
		}); err != nil {
			logrus.Errorf("api添加用户失败 apiPort: %s err: %v", apiPort, err)
			continue
		}
	}

	removeApiUserVos, err := service.SelectUsersPassword(false)
	if err != nil {
		logrus.Errorf("查询api删除用户失败 err: %v\n", err)
		return
	}
	for _, apiUserVo := range removeApiUserVos {
		if err := api.DeleteUser(apiUserVo.Password); err != nil {
			logrus.Errorf("api删除用户失败 apiPort: %s err: %v", apiPort, err)
			continue
		}
	}
}

func (t *TrojanGoProcess) handlerUserUploadAndDownload(apiPort string) {
	for {
		if !t.IsRunning(apiPort) {
			break
		}
	}
}

func (t *TrojanGoProcess) removeUsers(apiPort string) {
	for {
		if !t.IsRunning(apiPort) {
			break
		}
	}
}
