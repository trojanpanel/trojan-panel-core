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
		t := &TrojanGoProcess{process{binaryType: 2}}
		binaryFilePath, err := util.GetBinaryFile(2)
		if err != nil {
			return nil, err
		}
		configFilePath, err := util.GetConfigFile(2, apiPort)
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
			go t.handlerUserUploadAndDownload(apiPort)
			go t.handlerUsers(apiPort)
			return nil
		}
		logrus.Errorf("start trojan-go error err: process not found\n")
		return errors.New(constant.TrojanGoStartError)
	}
	logrus.Errorf("start trojan-go error err: lock not acquired\n")
	return errors.New(constant.TrojanGoStartError)
}

func (t *TrojanGoProcess) handlerUsers(apiPort string) {
	api := trojango.NewTrojanGoApi(apiPort)
	for {
		if !t.IsRunning(apiPort) {
			break
		}
		addApiUserVos, err := service.SelectUsersPassword(true)
		if err != nil {
			logrus.Errorf("数据库同步至Trojan Go apiPort: %s 查询用户失败 err: %v\n", apiPort, err)
		} else {
			for _, apiUserVo := range addApiUserVos {
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
					logrus.Errorf("数据库同步至Trojan Go apiPort: %s 添加用户失败 err: %v", apiPort, err)
					continue
				}
			}
		}
		apiUserVos, err := service.SelectUsersPassword(false)
		if err != nil {
			logrus.Errorf("数据库同步至Trojan Go apiPort: %s 查询用户失败 err: %v\n", apiPort, err)
		} else {
			for _, apiUser := range apiUserVos {
				if err := api.DeleteUser(apiUser.Password); err != nil {
					logrus.Errorf("数据库同步至Trojan Go apiPort: %s 删除用户失败 err: %v", apiPort, err)
					continue
				}
			}
		}
	}
}

func (t *TrojanGoProcess) handlerUserUploadAndDownload(apiPort string) {
	api := trojango.NewTrojanGoApi(apiPort)
	for {
		if !t.IsRunning(apiPort) {
			break
		}
		users, err := api.ListUsers()
		if err != nil {
			continue
		}
		for _, user := range users {
			downloadTraffic := int(user.GetTrafficTotal().GetDownloadTraffic())
			uploadTraffic := int(user.GetTrafficTotal().GetDownloadTraffic())
			if err := service.UpdateUser(user.GetUser().GetPassword(), &downloadTraffic,
				&uploadTraffic, nil); err != nil {
				logrus.Errorf("Trojan Go同步至数据库 apiPort: %s 更新用户失败 err: %v", apiPort, err)
				continue
			}
		}
	}
}
