package process

import (
	"errors"
	"os/exec"
	"runtime"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

type trojanGoProcess struct {
	process
}

func NewTrojanGoProcess(id int, apiPort string) (*xrayProcess, error) {
	x := &xrayProcess{
		process{
			apiPort: apiPort,
		},
	}
	x.apiPort = apiPort
	binaryFilePath, err := util.GetBinaryFilePath("trojan-go")
	if err != nil {
		return nil, errors.New("")
	}
	configFilePath, err := util.GetConfigPath(id, "trojan-go")
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(binaryFilePath, "-config", configFilePath)
	x.cmdMap.Store(id, cmd)
	runtime.SetFinalizer(x, x.Stop(id))
	return x, nil
}

func (x *xrayProcess) StartTrojanGo(id int) error {
	if x.IsRunning(id) {
		return nil
	}
	cmd, ok := x.cmdMap.Load(id)
	if ok {
		if err := cmd.(*exec.Cmd).Start(); err != nil {
			return errors.New(constant.TrojanGoStartError)
		}
	}
	return errors.New(constant.TrojanGoStartError)
}
