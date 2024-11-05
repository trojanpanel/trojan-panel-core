package process

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"trojan-core/model/constant"
	"trojan-core/util"
)

type HysteriaProcess struct {
	process
	proxy   string
	binPath string
}

var mutexHysteria sync.Mutex
var cmdMapHysteria sync.Map
var hysteriaProcess *HysteriaProcess

func init() {
	hysteriaProcess = &HysteriaProcess{
		process{mutex: &mutexHysteria, cmdMap: &cmdMapHysteria},
		constant.Hysteria,
		fmt.Sprintf("%s%s", constant.HysteriaBinPath, constant.Hysteria)}
}

func NewHysteriaInstance() *HysteriaProcess {
	return hysteriaProcess
}

func (h *HysteriaProcess) StartHysteria(apiPort uint) error {
	configPath := fmt.Sprintf("%s%d.json", constant.HysteriaPath, apiPort)
	if err := h.start(apiPort, h.binPath, "-c", configPath, "server"); err != nil {
		_ = util.RemoveFile(configPath)
		logrus.Errorf("start hysteria err: %v", err)
		return errors.New(constant.SysError)
	}
	return nil
}

func (h *HysteriaProcess) StopHysteria(apiPort uint) error {
	if err := h.stop(apiPort); err != nil {
		logrus.Errorf("stop hysteria err: %v", err)
		return errors.New(constant.SysError)
	}
	_ = util.RemoveFile(fmt.Sprintf("%s%d.json", constant.HysteriaPath, apiPort))
	return nil
}
