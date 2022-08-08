package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"sync"
	"trojan-panel-core/module/constant"
)

type process struct {
	mutex   sync.Mutex
	cmdMap  sync.Map[int, exec.Cmd]
	apiPort string
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
	defer p.mutex.Unlock()
	if p.mutex.TryLock() {
		if !p.IsRunning(id) {
			return nil
		}
		cmd, ok := p.cmdMap.Load(id)
		if ok {
			if err := cmd.(*exec.Cmd).Process.Kill(); err != nil {
				logrus.Errorf("stop process error id: %d err: %v\n", id, err)
				return errors.New(constant.ProcessStopError)
			}
			return nil
		}
		logrus.Errorf("stop process error id: %d err: process not found\n", id)
		return errors.New(constant.ProcessStopError)
	}
	logrus.Errorf("stop process error err: lock not acquired\n")
	return errors.New(constant.ProcessStopError)
}
