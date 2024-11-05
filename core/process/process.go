package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"sync"
	"trojan-core/model/constant"
)

type process struct {
	mutex  *sync.Mutex
	cmdMap *sync.Map
}

func (p *process) getCmdMap() *sync.Map {
	return p.cmdMap
}

func (p *process) isRunning(apiPort uint) bool {
	cmd, ok := p.cmdMap.Load(apiPort)
	if ok {
		return cmd != nil && cmd.(*exec.Cmd).Process != nil && cmd.(*exec.Cmd).ProcessState != nil
	}
	return false
}

func (p *process) start(apiPort uint, name string, arg ...string) error {
	if !p.mutex.TryLock() {
		logrus.Errorf("process start err: lock not acquired")
		return errors.New(constant.SysError)
	}
	defer p.mutex.Unlock()

	if p.isRunning(apiPort) {
		return nil
	}

	cmd := exec.Command(name, arg...)
	if cmd.Err != nil {
		logrus.Errorf("process err: %v", cmd.Err)
		return errors.New(constant.SysError)
	}

	if err := cmd.Start(); err != nil {
		logrus.Errorf("process start err: %v", err)
		return errors.New(constant.SysError)
	}

	p.cmdMap.Store(apiPort, cmd)
	return nil
}

func (p *process) stop(apiPort uint) error {
	if !p.mutex.TryLock() {
		logrus.Errorf("process stop err: lock not acquired")
		return errors.New(constant.SysError)
	}
	defer p.mutex.Unlock()

	if !p.isRunning(apiPort) {
		return nil
	}

	cmd, ok := p.cmdMap.Load(apiPort)
	if ok {
		if err := cmd.(*exec.Cmd).Process.Kill(); err != nil {
			logrus.Errorf("process stop err: %v", err)
			return errors.New(constant.SysError)
		}
		p.cmdMap.Delete(apiPort)
		return nil
	}
	logrus.Errorf("process stop error apiPort: %d err: process not found", apiPort)
	return errors.New(constant.SysError)
}

func (p *process) release(apiPort uint) error {
	if !p.mutex.TryLock() {
		logrus.Errorf("process release err: lock not acquired")
		return errors.New(constant.SysError)
	}
	defer p.mutex.Unlock()

	if !p.isRunning(apiPort) {
		return nil
	}

	cmd, ok := p.cmdMap.Load(apiPort)
	if ok {
		if err := cmd.(*exec.Cmd).Process.Release(); err != nil {
			logrus.Errorf("process release err: %v", err)
			return errors.New(constant.SysError)
		}
		p.cmdMap.Delete(apiPort)
		return nil
	}
	logrus.Errorf("process release error apiPort: %d err: process not found", apiPort)
	return errors.New(constant.SysError)
}
