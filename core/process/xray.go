package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"sync"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

type XrayProcess struct {
	process
}

func NewXrayProcess(apiPort string) (*XrayProcess, error) {
	var mutex sync.Mutex
	defer mutex.Unlock()
	if mutex.TryLock() {
		x := &XrayProcess{}
		binaryFilePath, err := util.GetBinaryFile("xray")
		if err != nil {
			return nil, err
		}
		configFilePath, err := util.GetConfigFile(apiPort, "xray")
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(binaryFilePath, "-c", configFilePath)
		x.cmdMap.Store(0, cmd)
		runtime.SetFinalizer(x, x.Stop(apiPort))
		return x, nil
	}
	logrus.Errorf("new xray process errror err: lock not acquired\n")
	return nil, errors.New(constant.NewXrayProcessError)
}

func (x *XrayProcess) StartXray(apiPort string) error {
	defer x.mutex.Unlock()
	if x.mutex.TryLock() {
		if x.IsRunning(apiPort) {
			return nil
		}
		cmd, ok := x.cmdMap.Load(apiPort)
		if ok {
			if err := cmd.(*exec.Cmd).Start(); err != nil {
				logrus.Errorf("start xray error err: %v\n", err)
				return errors.New(constant.XrayStartError)
			}
			return nil
		}
		logrus.Errorf("start xray error err: process not found\n")
		return errors.New(constant.XrayStartError)
	}
	logrus.Errorf("start xray error err: lock not acquired\n")
	return errors.New(constant.XrayStartError)
}
