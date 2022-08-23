package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"trojan-panel-core/module/constant"
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
		return nil
	}
	logrus.Errorf("start xray error err: lock not acquired\n")
	return errors.New(constant.XrayStartError)
}
