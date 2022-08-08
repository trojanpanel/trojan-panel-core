package process

import (
	"errors"
	"os/exec"
	"runtime"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

type xrayProcess struct {
	process
}

func NewXrayProcess(apiPort string) (*xrayProcess, error) {
	x := &xrayProcess{
		process{
			apiPort: apiPort,
		},
	}
	binaryFilePath, err := util.GetBinaryFilePath("xray")
	if err != nil {
		return nil, err
	}
	configFilePath, err := util.GetConfigPath(0, "xray")
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(binaryFilePath, "-c", configFilePath)
	x.cmdMap.Store(0, cmd)
	runtime.SetFinalizer(x, x.Stop(0))
	return x, nil
}

func (x *xrayProcess) StartXray(id int) error {
	if x.IsRunning(id) {
		return nil
	}
	cmd, ok := x.cmdMap.Load(id)
	if ok {
		if err := cmd.(*exec.Cmd).Start(); err != nil {
			return errors.New(constant.XrayStartError)
		}
	}
	return errors.New(constant.XrayStartError)
}
