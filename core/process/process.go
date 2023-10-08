package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"sync"
	"trojan-panel-core/model/constant"
	"trojan-panel-core/util"
)

type process struct {
	mutex      *sync.Mutex
	cmdMap     *sync.Map
	binaryType int // 1/xray 2/trojan-go 3/hysteria
}

func (p *process) GetCmdMap() *sync.Map {
	return p.cmdMap
}

func (p *process) IsRunning(apiPort uint) bool {
	cmd, ok := p.cmdMap.Load(apiPort)
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

func (p *process) Stop(apiPort uint, removeFile bool) error {
	defer p.mutex.Unlock()
	if p.mutex.TryLock() {
		if !p.IsRunning(apiPort) {
			logrus.Errorf("process has been stoped. apiPort: %d", apiPort)
			if removeFile {
				configFile, err := util.GetConfigFile(p.binaryType, apiPort)
				if err == nil {
					if err = util.RemoveFile(configFile); err != nil {
						return err
					}
				}
			}
			return nil
		}
		cmd, ok := p.cmdMap.Load(apiPort)
		if ok {
			if err := cmd.(*exec.Cmd).Process.Kill(); err != nil {
				logrus.Errorf("stop process error. apiPort: %d err: %v", apiPort, err)
				return errors.New(constant.ProcessStopError)
			}
			p.cmdMap.Delete(apiPort)
			if removeFile {
				configFile, err := util.GetConfigFile(p.binaryType, apiPort)
				if err == nil {
					if err = util.RemoveFile(configFile); err != nil {
						return err
					}
				}
			}
			return nil
		}
		logrus.Errorf("stop process error apiPort: %d err: process not found", apiPort)
		return errors.New(constant.ProcessStopError)
	}
	logrus.Errorf("stop process error err: lock not acquired")
	return errors.New(constant.ProcessStopError)
}

func GetState(nodeTypeId uint, apiPort uint) bool {
	switch nodeTypeId {
	case constant.Xray:
		return GetXrayState(apiPort)
	case constant.TrojanGo:
		return GetTrojanGoState(apiPort)
	case constant.Hysteria:
		return GetHysteriaState(apiPort)
	case constant.NaiveProxy:
		return GetNaiveProxyState(apiPort)
	case constant.Hysteria2:
		return GetHysteria2State(apiPort)
	default:
		return false
	}
}
