package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"sync"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

var mutexHysteria sync.Mutex
var cmdMapHysteria sync.Map

type HysteriaProcess struct {
	process
}

func NewHysteriaInstance() *HysteriaProcess {
	return &HysteriaProcess{process{mutex: &mutexHysteria, binaryType: 3, cmdMap: &cmdMapHysteria}}
}

func (h *HysteriaProcess) StopHysteriaInstance() error {
	apiPorts, err := util.GetConfigApiPorts(constant.HysteriaPath)
	if err != nil {
		return err
	}
	for _, apiPort := range apiPorts {
		if err = h.Stop(apiPort, true); err != nil {
			return err
		}
	}
	return nil
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
		if err := cmd.Start(); err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("start hysteria error err: %v", err)
			return errors.New(constant.HysteriaStartError)
		}
		return nil
	}
	logrus.Errorf("start hysteria error err: lock not acquired")
	return errors.New(constant.HysteriaStartError)
}
