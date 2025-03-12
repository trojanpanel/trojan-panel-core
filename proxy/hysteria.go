package proxy

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
	"trojan-core/model/constant"
)

var (
	hysteriaLogger logrus.Logger
	HysteriaCmdMap sync.Map
)

type HysteriaInstance struct {
	Instance
}

func init() {
	hysteriaLogger.SetOutput(&lumberjack.Logger{
		Filename:   constant.HysteriaLogPath,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	})
	hysteriaLogger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	hysteriaLogger.SetLevel(logrus.InfoLevel)
}

func NewHysteriaInstance(key string, configPath string) *HysteriaInstance {
	return &HysteriaInstance{
		Instance{
			BinPath:    constant.BinDir,
			Key:        key,
			ConfigPath: configPath,
			Command:    []string{"server", "-c", configPath},
			process: process{
				logger: &hysteriaLogger,
				cmdMap: &HysteriaCmdMap,
			},
		},
	}
}
