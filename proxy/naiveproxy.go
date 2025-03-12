package proxy

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
	"trojan-core/model/constant"
)

var (
	naiveproxyLogger logrus.Logger
	NaiveProxyCmdMap sync.Map
)

type NaiveProxyInstance struct {
	Instance
}

func init() {
	naiveproxyLogger.SetOutput(&lumberjack.Logger{
		Filename:   constant.NaiveProxyLogPath,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	})
	naiveproxyLogger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	naiveproxyLogger.SetLevel(logrus.InfoLevel)
}

func NewNaiveProxyInstance(key string, configPath string) *NaiveProxyInstance {
	return &NaiveProxyInstance{
		Instance{
			BinPath:    constant.BinDir,
			Key:        key,
			ConfigPath: configPath,
			Command:    []string{"run", "--config", configPath},
			process: process{
				logger: &naiveproxyLogger,
				cmdMap: &NaiveProxyCmdMap,
			},
		},
	}
}
