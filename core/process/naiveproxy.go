package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"sync"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

var mutexNaiveProxy sync.Mutex
var cmdMapNaiveProxy sync.Map

type NaiveProxyProcess struct {
	process
}

func NewNaiveProxyInstance() *NaiveProxyProcess {
	return &NaiveProxyProcess{process{mutex: &mutexNaiveProxy, binaryType: constant.NaiveProxy, cmdMap: &cmdMapNaiveProxy}}
}

func (n *NaiveProxyProcess) StopNaiveProxyInstance() error {
	apiPorts, err := util.GetConfigApiPorts(constant.NaiveProxyPath)
	if err != nil {
		return err
	}
	for _, apiPort := range apiPorts {
		if err = n.Stop(apiPort, true); err != nil {
			return err
		}
	}
	return nil
}

func (n *NaiveProxyProcess) StartNaiveProxy(apiPort uint) error {
	defer n.mutex.Unlock()
	if n.mutex.TryLock() {
		if n.IsRunning(apiPort) {
			return nil
		}
		binaryFilePath, err := util.GetBinaryFile(constant.NaiveProxy)
		if err != nil {
			return err
		}
		configFilePath, err := util.GetConfigFile(constant.NaiveProxy, apiPort)
		if err != nil {
			return err
		}
		cmd := exec.Command(binaryFilePath, "run", "--config", configFilePath)
		if cmd.Err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("naiveproxy command error err: %v", err)
			return errors.New(constant.NaiveProxyStartError)
		}
		if err := cmd.Start(); err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("start naiveproxy error err: %v", err)
			return errors.New(constant.NaiveProxyStartError)
		}
		n.cmdMap.Store(apiPort, cmd)
		go func() {
			if err := cmd.Wait(); err != nil {
				logrus.Errorf("naiveproxy process wait error err: %v", err)
			}
			if !cmd.ProcessState.Success() {
				logrus.Errorf("naiveproxy process state fail err: %v", err)
				if err := cmd.Process.Release(); err != nil {
					logrus.Errorf("naiveproxy process release error err: %v", err)
					return
				}
				if err = util.RemoveFile(configFilePath); err != nil {
					logrus.Errorf("naiveproxy process remove file error err: %v", err)
					return
				}
				n.cmdMap.Delete(apiPort)
			}
		}()
		return nil
	}
	logrus.Errorf("start naiveproxy error err: lock not acquired")
	return errors.New(constant.NaiveProxyStartError)
}

func GetNaiveProxyState(apiPort uint) bool {
	_, ok := NewNaiveProxyInstance().GetCmdMap().Load(apiPort)
	return ok
}
