package process

import (
	"errors"
	"os/exec"
	"sync"
	"trojan-panel-core/module/constant"
)

type process struct {
	cmdMap  sync.Map[int, exec.Cmd]
	apiPort string
	name    string
}

func (p *process) IsRunning(id int) bool {
	cmd, ok := p.cmdMap.Load(id)
	if ok {
		if cmd == nil || cmd.(*exec.Cmd).Process == nil {
			return false
		}
		if cmd.(*exec.Cmd).ProcessState == nil {
			return true
		}
	}
	return false
}

func (p *process) Stop(id int) error {
	if !p.IsRunning(id) {
		return nil
	}
	cmd, ok := p.cmdMap.Load(id)
	if ok {
		if err := cmd.(*exec.Cmd).Process.Kill(); err != nil {
			return errors.New(constant.ProcessStopError)
		}
	}
	return errors.New(constant.ProcessStopError)
}
