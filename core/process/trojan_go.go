package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

type TrojanGoProcess struct {
	process
}

func NewTrojanGoInstance() *TrojanGoProcess {
	t := &TrojanGoProcess{process{mutex: &mutex, binaryType: 2, cmdMap: &cmdMap}}
	runtime.SetFinalizer(t, t.StopTrojanGoInstance())
	return t
}

func (t *TrojanGoProcess) StopTrojanGoInstance() error {
	apiPorts, err := util.GetConfigApiPorts(constant.TrojanGoPath)
	if err != nil {
		return err
	}
	for _, apiPort := range apiPorts {
		if err = t.Stop(apiPort, true); err != nil {
			return err
		}
	}
	return nil
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
		if err := cmd.Start(); err != nil {
			logrus.Errorf("start trojan-go error err: %v\n", err)
			return errors.New(constant.TrojanGoStartError)
		}
		return nil
	}
	logrus.Errorf("start trojan-go error err: lock not acquired\n")
	return errors.New(constant.TrojanGoStartError)
}
