package proxy

import (
	"github.com/avast/retry-go"
	"github.com/sirupsen/logrus"
	"time"
	"trojan-core/util"
)

type Instance struct {
	BinPath    string
	Key        string
	ConfigPath string
	Command    []string
	process
}

func (i *Instance) IsRunning() bool {
	return i.isRunning(i.Key)
}

func (i *Instance) Start() error {
	if err := i.start(i.Key, i.BinPath, i.Command...); err != nil {
		tryRemoveFile(i.ConfigPath)
		logrus.Errorf("start bingPath: %s err: %v", i.BinPath, err)
		return err
	}
	return nil
}

func (i *Instance) Stop() error {
	if err := i.stop(i.Key); err != nil {
		logrus.Errorf("stop bingPath: %s err: %v", i.BinPath, err)
		return err
	}
	tryRemoveFile(i.ConfigPath)
	return nil
}

func tryRemoveFile(filePath string) {
	if err := retry.Do(func() error {
		return util.RemoveFile(filePath)
	}, []retry.Option{
		retry.Delay(3 * time.Second),
		retry.Attempts(2),
	}...); err != nil {
		logrus.Errorf(err.Error())
	}
}
