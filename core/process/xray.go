package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"sync"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

var mutexXray sync.Mutex
var cmdMapXray sync.Map

type XrayProcess struct {
	process
}

func NewXrayProcess() *XrayProcess {
	return &XrayProcess{process{mutex: &mutexXray, binaryType: constant.Xray, cmdMap: &cmdMapXray}}
}

func (x *XrayProcess) StopXrayProcess() error {
	apiPorts, err := util.GetConfigApiPorts(constant.XrayPath)
	if err != nil {
		return err
	}
	for _, apiPort := range apiPorts {
		if err = x.Stop(apiPort, true); err != nil {
			return err
		}
	}
	return nil
}

func (x *XrayProcess) StartXray(apiPort uint) error {
	defer x.mutex.Unlock()
	if x.mutex.TryLock() {
		if x.IsRunning(apiPort) {
			return nil
		}
		binaryFilePath, err := util.GetBinaryFile(constant.Xray)
		if err != nil {
			return err
		}
		configFilePath, err := util.GetConfigFile(constant.Xray, apiPort)
		if err != nil {
			return err
		}
		cmd := exec.Command(binaryFilePath, "-c", configFilePath)
		if cmd.Err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("xray command error err: %v", err)
			return errors.New(constant.XrayStartError)
		}
		if err := cmd.Start(); err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("start xray error err: %v", err)
			return errors.New(constant.XrayStartError)
		}
		x.cmdMap.Store(apiPort, cmd)
		go func() {
			if err := cmd.Wait(); err != nil {
				logrus.Errorf("xray process wait error err: %v", err)
			}
			if !cmd.ProcessState.Success() {
				logrus.Errorf("xray process state fail err: %v", err)
				if err := cmd.Process.Release(); err != nil {
					logrus.Errorf("xray process release error err: %v", err)
					return
				}
				if err = util.RemoveFile(configFilePath); err != nil {
					logrus.Errorf("xray process remove file error err: %v", err)
					return
				}
				x.cmdMap.Delete(apiPort)
			}
		}()
		return nil
	}
	logrus.Errorf("start xray error err: lock not acquired")
	return errors.New(constant.XrayStartError)
}

func GetXrayState(apiPort uint) bool {
	_, ok := NewXrayProcess().GetCmdMap().Load(apiPort)
	return ok
}
