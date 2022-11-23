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

func (t *NaiveProxyProcess) StopNaiveProxyInstance() error {
	apiPorts, err := util.GetConfigApiPorts(constant.NaiveProxyPath)
	if err != nil {
		return err
	}
	for _, apiPort := range apiPorts {
		if err = t.Stop(apiPort, true); err != nil {
			return err
		}
	}
	return nil
}

func (t *NaiveProxyProcess) StartNaiveProxy(apiPort uint) error {
	defer t.mutex.Unlock()
	if t.mutex.TryLock() {
		if t.IsRunning(apiPort) {
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
		cmd := exec.Command(binaryFilePath, "run --config", configFilePath)
		t.cmdMap.Store(apiPort, cmd)
		if err := cmd.Start(); err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("start naiveproxy error err: %v", err)
			return errors.New(constant.NaiveProxyStartError)
		}
		return nil
	}
	logrus.Errorf("start naiveproxy error err: lock not acquired")
	return errors.New(constant.NaiveProxyStartError)
}
