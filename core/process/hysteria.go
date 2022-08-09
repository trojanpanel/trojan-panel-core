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

type HysteriaProcess struct {
	process
}

func NewHysteriaProcess(apiPort string) (*HysteriaProcess, error) {
	var mutex sync.Mutex
	defer mutex.Unlock()
	if mutex.TryLock() {
		h := &HysteriaProcess{process{binaryType: 3}}
		binaryFilePath, err := util.GetBinaryFile(3)
		if err != nil {
			return nil, err
		}
		configFilePath, err := util.GetConfigFile(3, apiPort)
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(binaryFilePath, "-c", configFilePath, "server")
		h.cmdMap.Store(apiPort, cmd)
		runtime.SetFinalizer(h, h.Stop(apiPort))
		return h, nil
	}
	logrus.Errorf("new hysteria process errror err: lock not acquired\n")
	return nil, errors.New(constant.NewHysteriaProcessError)
}

func (h *HysteriaProcess) StartHysteria(apiPort string) error {
	defer h.mutex.Unlock()
	if h.mutex.TryLock() {
		if h.IsRunning(apiPort) {
			return nil
		}
		cmd, ok := h.cmdMap.Load(apiPort)
		if ok {
			if err := cmd.(*exec.Cmd).Start(); err != nil {
				logrus.Errorf("start hysteria error err: %v\n", err)
				return errors.New(constant.HysteriaStartError)
			}
			return nil
		}
		logrus.Errorf("start hysteria error err: process not found\n")
		return errors.New(constant.HysteriaStartError)
	}
	logrus.Errorf("start hysteria error err: lock not acquired\n")
	return errors.New(constant.HysteriaStartError)
}
