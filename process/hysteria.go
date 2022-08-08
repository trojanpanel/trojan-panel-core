package process

import (
	"errors"
	"os/exec"
	"runtime"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

type hysteriaProcess struct {
	process
}

func NewHysteriaProcess(id int, apiPort string) (*hysteriaProcess, error) {
	h := &hysteriaProcess{
		process{
			apiPort: apiPort,
		},
	}
	binaryFilePath, err := util.GetBinaryFilePath("hysteria")
	if err != nil {
		return nil, err
	}
	configFilePath, err := util.GetConfigPath(id, "hysteria")
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(binaryFilePath, "-c", configFilePath, "server")
	h.cmdMap.Store(id, cmd)
	runtime.SetFinalizer(h, h.Stop(id))
	return h, nil
}

func (h *hysteriaProcess) StartHysteria(id int) error {
	if h.IsRunning(id) {
		return nil
	}
	cmd, ok := h.cmdMap.Load(id)
	if ok {
		if err := cmd.(*exec.Cmd).Start(); err != nil {
			return errors.New(constant.HysteriaStartError)
		}
	}
	return errors.New(constant.HysteriaStartError)
}
