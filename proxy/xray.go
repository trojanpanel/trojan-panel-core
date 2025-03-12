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
	xrayLogger logrus.Logger
	XrayCmdMap sync.Map
)

type XrayInstance struct {
	Instance
}

func init() {
	xrayLogger.SetOutput(&lumberjack.Logger{
		Filename:   constant.XrayLogPath,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	})
	xrayLogger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	xrayLogger.SetLevel(logrus.InfoLevel)
}

func NewXrayInstance(key string, configPath string) *XrayInstance {
	return &XrayInstance{
		Instance{
			BinPath:    constant.BinDir,
			Key:        key,
			ConfigPath: configPath,
			Command:    []string{"run", "-c", configPath},
			process: process{
				logger: &xrayLogger,
				cmdMap: &XrayCmdMap,
			},
		},
	}
}

func GetXrayBinPath() string {
	return constant.BinDir + GetXrayBinName()
}

func GetXrayBinName() string {
	xrayFileName := fmt.Sprintf("xray-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		xrayFileName += ".exe"
	}
	return xrayFileName
}

func DownloadXray(version string) error {
	return util.DownloadFromGithub(GetXrayBinName(), GetXrayBinPath(), "XTLS", "Xray-core", version)
}
