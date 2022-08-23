package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"trojan-panel-core/app/xray"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/service"
	"trojan-panel-core/util"
)

type XrayProcess struct {
	process
}

func NewXrayProcess() *XrayProcess {
	return &XrayProcess{process{mutex: &mutex, binaryType: 1, cmdMap: &cmdMap}}
}

func (x *XrayProcess) StartXray(apiPort uint) error {
	defer x.mutex.Unlock()
	if x.mutex.TryLock() {
		if x.IsRunning(apiPort) {
			return nil
		}
		binaryFilePath, err := util.GetBinaryFile(1)
		if err != nil {
			return err
		}
		configFilePath, err := util.GetConfigFile(1, apiPort)
		if err != nil {
			return err
		}
		cmd := exec.Command(binaryFilePath, "-c", configFilePath)
		x.cmdMap.Store(0, cmd)
		runtime.SetFinalizer(x, x.Stop(apiPort))
		if err := cmd.Start(); err != nil {
			logrus.Errorf("start xray error err: %v\n", err)
			return errors.New(constant.XrayStartError)
		}
		go x.handlerUserUploadAndDownload(apiPort)
		return nil
	}
	logrus.Errorf("start xray error err: lock not acquired\n")
	return errors.New(constant.XrayStartError)
}

// 将数据库中的用户同步至应用
func (x *XrayProcess) handlerUserUploadAndDownload(apiPort uint) {
	protocol, err := util.GetXrayProtocolByApiPort(apiPort)
	if err != nil {
		logrus.Errorf("数据库同步至Xray apiPort: %d 未查询到xray的协议 err: %v\n", apiPort, err)
		return
	}
	api := xray.NewXrayApi(apiPort)
	for {
		if !x.IsRunning(apiPort) {
			logrus.Errorf("数据库同步至Xray apiPort: %d xray not running\n", apiPort)
			break
		}

		addPasswords, err := service.SelectAccountPasswords(false)
		if err != nil {
			logrus.Errorf("数据库同步至Xray apiPort: %d 查询用户失败 err: %v\n", apiPort, err)
		} else {
			for _, item := range addPasswords {
				// 如果应用中存在则跳过
				stats, err := api.GetUserStats(item, "downlink", true)
				if err != nil || stats != nil {
					continue
				}
				userDto := dto.XrayAddUserDto{
					Protocol: protocol,
					Password: item,
				}
				if err := api.AddUser(userDto); err != nil {
					logrus.Errorf("数据库同步至Xray apiPort: %d 添加用户失败 err: %v", apiPort, err)
					continue
				}
			}
		}
	}
}
