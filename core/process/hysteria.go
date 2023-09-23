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

var mutexHysteria sync.Mutex
var cmdMapHysteria sync.Map

type HysteriaProcess struct {
	process
}

func NewHysteriaInstance() *HysteriaProcess {
	return &HysteriaProcess{process{mutex: &mutexHysteria, binaryType: constant.Hysteria, cmdMap: &cmdMapHysteria}}
}

func (h *HysteriaProcess) StopHysteriaInstance() error {
	apiPorts, err := util.GetConfigApiPorts(constant.HysteriaPath)
	if err != nil {
		return err
	}
	for _, apiPort := range apiPorts {
		if err = h.Stop(apiPort, true); err != nil {
			return err
		}
	}
	return nil
}

func (h *HysteriaProcess) StartHysteria(apiPort uint) error {
	defer h.mutex.Unlock()
	if h.mutex.TryLock() {
		if h.IsRunning(apiPort) {
			return nil
		}
		binaryFilePath, err := util.GetBinaryFile(constant.Hysteria)
		if err != nil {
			return err
		}
		configFilePath, err := util.GetConfigFile(constant.Hysteria, apiPort)
		if err != nil {
			return err
		}
		cmd := exec.Command(binaryFilePath, "-c", configFilePath, "server")
		if cmd.Err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("hysteria command error err: %v", err)
			return errors.New(constant.HysteriaStartError)
		}
		if err := cmd.Start(); err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("start hysteria error err: %v", err)
			return errors.New(constant.HysteriaStartError)
		}
		h.cmdMap.Store(apiPort, cmd)

		// timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		done := make(chan error)
		go func() {
			done <- cmd.Wait()
			select {
			case err := <-done:
				if err != nil {
					logrus.Errorf("hysteria process wait error err: %v", err)
					h.releaseProcess(apiPort, configFilePath)
				}
			case <-ctx.Done():
				logrus.Errorf("hysteria process wait timeout err: %v", err)
				h.releaseProcess(apiPort, configFilePath)
			}
		}()
		return nil
	}
	logrus.Errorf("start hysteria error err: lock not acquired")
	return errors.New(constant.HysteriaStartError)
}

func (h *HysteriaProcess) releaseProcess(apiPort uint, configFilePath string) {
	load, ok := NewHysteriaInstance().GetCmdMap().Load(apiPort)
	if ok {
		cmd := load.(*exec.Cmd)
		if !cmd.ProcessState.Success() {
			h.cmdMap.Delete(apiPort)
			if err := cmd.Process.Release(); err != nil {
				logrus.Errorf("hysteria process release error err: %v", err)
			}
			if err := util.RemoveFile(configFilePath); err != nil {
				logrus.Errorf("hysteria process remove file error err: %v", err)
			}
		}
	}
}

func GetHysteriaState(apiPort uint) bool {
	_, ok := NewHysteriaInstance().GetCmdMap().Load(apiPort)
	return ok
}
