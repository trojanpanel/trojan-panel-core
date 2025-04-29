package proxy

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
	"sync"
	"trojan-core/model/constant"
	"trojan-core/util"
)

var (
	naiveProxyLogger logrus.Logger
	NaiveProxyCmdMap sync.Map
)

type NaiveProxyInstance struct {
	Instance
}

func init() {
	naiveProxyLogger.SetOutput(&lumberjack.Logger{
		Filename:   constant.NaiveProxyLogPath,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	})
	naiveProxyLogger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	naiveProxyLogger.SetLevel(logrus.InfoLevel)
}

func NewNaiveProxyInstance(key string) *NaiveProxyInstance {
	configPath := GetNaiveProxyConfigPath(key)
	return &NaiveProxyInstance{
		Instance{
			BinPath:    GetNaiveProxyBinPath(),
			Key:        key,
			ConfigPath: configPath,
			Command:    []string{"run", "--config", configPath},
			process: process{
				logger: &naiveProxyLogger,
				cmdMap: &NaiveProxyCmdMap,
			},
		},
	}
}

func GetNaiveProxyBinPath() string {
	return constant.BinDir + GetNaiveProxyBinName()
}

func GetNaiveProxyBinName() string {
	naiveProxyFileName := fmt.Sprintf("naiveproxy-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		naiveProxyFileName += ".exe"
	}
	return naiveProxyFileName
}

func GetNaiveProxyConfigPath(key string) string {
	return constant.NaiveProxyConfigDir + key + constant.NaiveProxyConfigExt
}

func DownloadNaiveProxy(version string) error {
	naiveProxyFileName := fmt.Sprintf("naive-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		naiveProxyFileName = naiveProxyFileName + ".exe"
	}
	naiveProxyBinPath := constant.BinDir + naiveProxyFileName
	if err := util.DownloadFromGithub(naiveProxyFileName, naiveProxyBinPath, "jonssonyan", "naive", version); err != nil {
		return err
	}
	if !util.Exists(naiveProxyBinPath) {
		return fmt.Errorf("naive bin file not exists")
	}
	if err := os.Rename(naiveProxyBinPath, GetNaiveProxyBinPath()); err != nil {
		return err
	}
	return nil
}
