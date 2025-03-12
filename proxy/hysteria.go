package proxy

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"runtime"
	"sync"
	"trojan-core/model/constant"
	"trojan-core/util"
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

func NewHysteriaInstance(key string) *HysteriaInstance {
	configPath := constant.HysteriaConfigDir + key + constant.HysteriaConfigExt
	return &HysteriaInstance{
		Instance{
			BinPath:    GetHysteriaBinPath(),
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

func GetHysteriaBinPath() string {
	return constant.BinDir + GetHysteriaBinName()
}

func GetHysteriaBinName() string {
	hysteriaFileName := fmt.Sprintf("hysteria-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		hysteriaFileName += ".exe"
	}
	return hysteriaFileName
}

func DownloadHysteria(version string) error {
	return util.DownloadFromGithub(GetHysteriaBinName(), GetHysteriaBinPath(), "apernet", "hysteria", version)
}
