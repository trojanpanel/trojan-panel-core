package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

type HysteriaProcess struct {
	process
}

func NewHysteriaInstance() *HysteriaProcess {
	return &HysteriaProcess{process{mutex: &mutex, binaryType: 3, cmdMap: &cmdMap}}
}

func (h *HysteriaProcess) StartHysteria(apiPort uint) error {
	defer h.mutex.Unlock()
	if h.mutex.TryLock() {
		if h.IsRunning(apiPort) {
			return nil
		}
		binaryFilePath, err := util.GetBinaryFile(3)
		if err != nil {
			return err
		}
		configFilePath, err := util.GetConfigFile(3, apiPort)
		if err != nil {
			return err
		}
		cmd := exec.Command(binaryFilePath, "-c", configFilePath, "server")
		h.cmdMap.Store(apiPort, cmd)
		runtime.SetFinalizer(h, h.Stop(apiPort))
		if err := cmd.Start(); err != nil {
			logrus.Errorf("start hysteria error err: %v\n", err)
			return errors.New(constant.HysteriaStartError)
		}
		return nil
	}
	logrus.Errorf("start hysteria error err: lock not acquired\n")
	return errors.New(constant.HysteriaStartError)
}
