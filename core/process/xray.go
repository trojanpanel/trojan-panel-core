package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"sync"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

var mutexXray sync.Mutex
var cmdMapXray sync.Map

type XrayProcess struct {
	process
}

func NewXrayProcess() *XrayProcess {
	return &XrayProcess{process{mutex: &mutexXray, binaryType: 1, cmdMap: &cmdMapXray}}
}

func (x *XrayProcess) StopXrayProcess() error {
	apiPorts, err := util.GetConfigApiPorts(constant.XrayPath)
	if err != nil {
		return err
	}
	for _, apiPort := range apiPorts {
		if err = x.Stop(apiPort, true); err != nil {
			return err
		}
	}
	return nil
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
		x.cmdMap.Store(apiPort, cmd)
		if err := cmd.Start(); err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("start xray error err: %v", err)
			return errors.New(constant.XrayStartError)
		}
		return nil
	}
	logrus.Errorf("start xray error err: lock not acquired")
	return errors.New(constant.XrayStartError)
}
