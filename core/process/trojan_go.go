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
		runtime.SetFinalizer(t, t.Stop(apiPort, true))
		if err := cmd.Start(); err != nil {
			logrus.Errorf("start trojan-go error err: %v\n", err)
			return errors.New(constant.TrojanGoStartError)
		}
		return nil
	}
	logrus.Errorf("start trojan-go error err: lock not acquired\n")
	return errors.New(constant.TrojanGoStartError)
}
