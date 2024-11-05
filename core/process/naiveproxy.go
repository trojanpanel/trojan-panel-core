package process

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"trojan-core/model/constant"
	"trojan-core/util"
)

type NaiveProcess struct {
	process
	proxy   string
	binPath string
}

var mutexNaive sync.Mutex
var cmdMapNaive sync.Map
var naiveProcess *NaiveProcess

func init() {
	naiveProcess = &NaiveProcess{
		process{mutex: &mutexNaive, cmdMap: &cmdMapNaive},
		constant.Naive,
		fmt.Sprintf("%s%s", constant.NaiveBinPath, constant.NaivePath)}
}

func NewNaiveProxyInstance() *NaiveProcess {
	return naiveProcess
}

func (h *NaiveProcess) StartNaive(apiPort uint) error {
	configPath := fmt.Sprintf("%s%d.json", constant.NaivePath, apiPort)
	if err := h.start(apiPort, h.binPath, "run", "--config", configPath); err != nil {
		_ = util.RemoveFile(configPath)
		logrus.Errorf("start naive err: %v", err)
		return errors.New(constant.SysError)
	}
	return nil
}

func (h *NaiveProcess) StopNaive(apiPort uint) error {
	if err := h.stop(apiPort); err != nil {
		logrus.Errorf("stop niave err: %v", err)
		return errors.New(constant.SysError)
	}
	_ = util.RemoveFile(fmt.Sprintf("%s%d.json", constant.NaivePath, apiPort))
	return nil
}
