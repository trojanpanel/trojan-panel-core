package proxy

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"runtime"
	"sync"
	"trojan-core/model/constant"
)

var (
	singBoxLogger logrus.Logger
	SingBoxCmdMap sync.Map
)

type SingBoxInstance struct {
	Instance
}

func init() {
	singBoxLogger.SetOutput(&lumberjack.Logger{
		Filename:   constant.SingBoxLogPath,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	})
	singBoxLogger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	singBoxLogger.SetLevel(logrus.InfoLevel)
}

func NewSingBoxInstance(key string) *SingBoxInstance {
	configPath := GetSingBoxConfigPath(key)
	return &SingBoxInstance{
		Instance{
			BinPath:    GetSingBoxBinPath(),
			Key:        key,
			ConfigPath: configPath,
			Command:    []string{"run", "-c", configPath},
			process: process{
				logger: &singBoxLogger,
				cmdMap: &SingBoxCmdMap,
			},
		},
	}
}

func GetSingBoxBinPath() string {
	return constant.BinDir + GetSingBoxBinName()
}

func GetSingBoxBinName() string {
	singBoxFileName := fmt.Sprintf("singbox-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		singBoxFileName += ".exe"
	}
	return singBoxFileName
}

func GetSingBoxConfigPath(key string) string {
	return constant.SingBoxConfigDir + key + constant.SingBoxConfigExt
}

func DownloadSingBox(version string) error {
	return nil
}
