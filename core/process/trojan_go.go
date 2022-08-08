package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"sync"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

type TrojanGoProcess struct {
	process
}

func NewTrojanGoProcess(id int, apiPort string) (*TrojanGoProcess, error) {
	var mutex sync.Mutex
	defer mutex.Unlock()
	if mutex.TryLock() {
		t := &TrojanGoProcess{
			process{
				apiPort: apiPort,
			},
		}
		t.apiPort = apiPort
		binaryFilePath, err := util.GetBinaryFile("trojan-go")
		if err != nil {
			return nil, err
		}
		configFilePath, err := util.GetConfigFile(id, "trojan-go")
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(binaryFilePath, "-config", configFilePath)
		t.cmdMap.Store(id, cmd)
		runtime.SetFinalizer(t, t.Stop(id))
		return t, nil
	}
	logrus.Errorf("new trojan-go process errror err: lock not acquired\n")
	return nil, errors.New(constant.NewTrojanGoProcessError)
}

func (t *TrojanGoProcess) StartTrojanGo(id int) error {
	defer t.mutex.Unlock()
	if t.mutex.TryLock() {
		if t.IsRunning(id) {
			return nil
		}
		cmd, ok := t.cmdMap.Load(id)
		if ok {
			if err := cmd.(*exec.Cmd).Start(); err != nil {
				logrus.Errorf("start trojan-go error err: %v\n", err)
				return errors.New(constant.TrojanGoStartError)
			}
			return nil
		}
		logrus.Errorf("start trojan-go error err: process not found\n")
		return errors.New(constant.TrojanGoStartError)
	}
	logrus.Errorf("start trojan-go error err: lock not acquired\n")
	return errors.New(constant.TrojanGoStartError)
}
