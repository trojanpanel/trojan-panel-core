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
		x := &XrayProcess{
			process{
				ApiPort: apiPort,
			},
		}
		binaryFilePath, err := util.GetBinaryFile("xray")
		if err != nil {
			return nil, err
		}
		configFilePath, err := util.GetConfigFile(0, "xray")
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(binaryFilePath, "-c", configFilePath)
		x.cmdMap.Store(0, cmd)
		runtime.SetFinalizer(x, x.Stop(0))
		return x, nil
	}
	logrus.Errorf("new xray process errror err: lock not acquired\n")
	return nil, errors.New(constant.NewXrayProcessError)
}

func (x *XrayProcess) StartXray(id int) error {
	defer x.mutex.Unlock()
	if x.mutex.TryLock() {
		if x.IsRunning(id) {
			return nil
		}
		cmd, ok := x.cmdMap.Load(id)
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
