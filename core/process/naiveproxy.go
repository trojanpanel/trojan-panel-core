package process

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"sync"
	"time"
	"trojan-panel-core/model/constant"
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

		// timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		done := make(chan error)
		go func() {
			done <- cmd.Wait()
			select {
			case err := <-done:
				if err != nil {
					logrus.Errorf("naiveproxy process wait error err: %v", err)
					n.releaseProcess(apiPort, configFilePath)
				}
			case <-ctx.Done():
				logrus.Errorf("naiveproxy process wait timeout err: %v", err)
				n.releaseProcess(apiPort, configFilePath)
			}
		}()
		return nil
	}
	logrus.Errorf("start naiveproxy error err: lock not acquired")
	return errors.New(constant.NaiveProxyStartError)
}

func (n *NaiveProxyProcess) releaseProcess(apiPort uint, configFilePath string) {
	load, ok := NewNaiveProxyInstance().GetCmdMap().Load(apiPort)
	if ok {
		cmd := load.(*exec.Cmd)
		if !cmd.ProcessState.Success() {
			n.cmdMap.Delete(apiPort)
			if err := cmd.Process.Release(); err != nil {
				logrus.Errorf("naiveproxy process release error err: %v", err)
			}
			if err := util.RemoveFile(configFilePath); err != nil {
				logrus.Errorf("naiveproxy process remove file error err: %v", err)
			}
		}
	}
}

func GetNaiveProxyState(apiPort uint) bool {
	_, ok := NewNaiveProxyInstance().GetCmdMap().Load(apiPort)
	return ok
}
