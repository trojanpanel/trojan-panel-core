package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"trojan-panel-core/app/trojango"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
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
		go t.handlerUsers(apiPort)
		return nil
	}
	logrus.Errorf("start trojan-go error err: lock not acquired\n")
	return errors.New(constant.TrojanGoStartError)
}

// 更新应用中的用户
func (t *TrojanGoProcess) handlerUsers(apiPort uint) {
	api := trojango.NewTrojanGoApi(apiPort)
	// 更新每个应用中的数据
	for {
		if !t.IsRunning(apiPort) {
			logrus.Errorf("数据库同步至Trojan Go apiPort: %d trojan go not running\n", apiPort)
			break
		}
		addUserApiVos, err := service.SelectUsersToApi(true)
		if err != nil {
			logrus.Errorf("数据库同步至Trojan Go apiPort: %d 查询用户失败 err: %v\n", apiPort, err)
		} else {
			for _, apiUserVo := range addUserApiVos {
				userStatus, err := api.GetUser(apiUserVo.Password)
				if err != nil || userStatus == nil || userStatus.GetUser() == nil {
					continue
				}
				if err := api.AddUser(dto.TrojanGoAddUserDto{
					Password:           apiUserVo.Password,
					UploadTraffic:      apiUserVo.Upload,
					DownloadTraffic:    apiUserVo.Download,
					IpLimit:            0,
					DownloadSpeedLimit: 0,
					UploadSpeedLimit:   0,
				}); err != nil {
					logrus.Errorf("数据库同步至Trojan Go apiPort: %d 添加用户失败 err: %v", apiPort, err)
					continue
				}
			}
		}
		apiUserVos, err := service.SelectUsersToApi(false)
		if err != nil {
			logrus.Errorf("数据库同步至Trojan Go apiPort: %d 查询用户失败 err: %v\n", apiPort, err)
		} else {
			for _, apiUser := range apiUserVos {
				if err := api.DeleteUser(apiUser.Password); err != nil {
					logrus.Errorf("数据库同步至Trojan Go apiPort: %d 删除用户失败 err: %v", apiPort, err)
					continue
				}
			}
		}
	}
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
			encodePassword, err := util.AesEncode(user.GetUser().GetPassword())
			if err != nil {
				continue
			}
			if err := service.UpdateUser(nil, &apiPort, &encodePassword, &downloadTraffic,
				&uploadTraffic); err != nil {
				logrus.Errorf("Trojan Go同步至数据库 apiPort: %d 更新用户失败 err: %v", apiPort, err)
				continue
			}
		}
	}
}
