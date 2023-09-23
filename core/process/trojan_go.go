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

var mutexTrojanGo sync.Mutex
var cmdMapTrojanGo sync.Map

type TrojanGoProcess struct {
	process
}

func NewTrojanGoInstance() *TrojanGoProcess {
	return &TrojanGoProcess{process{mutex: &mutexTrojanGo, binaryType: constant.TrojanGo, cmdMap: &cmdMapTrojanGo}}
}

func (t *TrojanGoProcess) StopTrojanGoInstance() error {
	apiPorts, err := util.GetConfigApiPorts(constant.TrojanGoPath)
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

func (t *TrojanGoProcess) StartTrojanGo(apiPort uint) error {
	defer t.mutex.Unlock()
	if t.mutex.TryLock() {
		if t.IsRunning(apiPort) {
			return nil
		}
		binaryFilePath, err := util.GetBinaryFile(constant.TrojanGo)
		if err != nil {
			return err
		}
		configFilePath, err := util.GetConfigFile(constant.TrojanGo, apiPort)
		if err != nil {
			return err
		}
		cmd := exec.Command(binaryFilePath, "-config", configFilePath)
		if cmd.Err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("trojan-go command error err: %v", err)
			return errors.New(constant.TrojanGoStartError)
		}
		if err := cmd.Start(); err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("start trojan-go error err: %v", err)
			return errors.New(constant.TrojanGoStartError)
		}
		t.cmdMap.Store(apiPort, cmd)

		// timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		done := make(chan error)
		go func() {
			done <- cmd.Wait()
			select {
			case err := <-done:
				if err != nil {
					logrus.Errorf("trojan-go process wait error err: %v", err)
					t.releaseProcess(apiPort, configFilePath)
				}
			case <-ctx.Done():
				logrus.Errorf("trojan-go process wait timeout err: %v", err)
				t.releaseProcess(apiPort, configFilePath)
			}
		}()
		return nil
	}
	logrus.Errorf("start trojan-go error err: lock not acquired")
	return errors.New(constant.TrojanGoStartError)
}

func (t *TrojanGoProcess) releaseProcess(apiPort uint, configFilePath string) {
	load, ok := NewTrojanGoInstance().GetCmdMap().Load(apiPort)
	if ok {
		cmd := load.(*exec.Cmd)
		if !cmd.ProcessState.Success() {
			t.cmdMap.Delete(apiPort)
			if err := cmd.Process.Release(); err != nil {
				logrus.Errorf("trojan-go process release error err: %v", err)
			}
			if err := util.RemoveFile(configFilePath); err != nil {
				logrus.Errorf("trojan-go process remove file error err: %v", err)
			}
		}
	}
}

func GetTrojanGoState(apiPort uint) bool {
	_, ok := NewTrojanGoInstance().GetCmdMap().Load(apiPort)
	return ok
}
