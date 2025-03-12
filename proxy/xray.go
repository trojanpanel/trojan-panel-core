package proxy

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

func NewXrayInstance(key string) *XrayInstance {
	configPath := GetXrayConfigPath(key)
	return &XrayInstance{
		Instance{
			BinPath:    GetXrayBinPath(),
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

func GetXrayConfigPath(key string) string {
	return constant.XrayConfigDir + key + constant.XrayConfigExt
}

func DownloadXray(version string) error {
	goarch := runtime.GOARCH
	if strings.Contains(goarch, "amd") {
		goarch = strings.TrimPrefix(goarch, "amd")
	}
	xrayFileName := fmt.Sprintf("Xray-%s-%s.zip", runtime.GOOS, goarch)
	xrayBinPath := constant.BinDir + xrayFileName
	if err := util.DownloadFromGithub(xrayFileName, xrayBinPath, "XTLS", "Xray-core", version); err != nil {
		return err
	}
	if err := util.Unzip(xrayBinPath, constant.BinDir); err != nil {
		return err
	}
	if !util.Exists(constant.BinDir + "xray") {
		return fmt.Errorf("xray bin file not exists")
	}
	if err := os.Rename(constant.BinDir+"xray", GetXrayBinPath()); err != nil {
		return err
	}
	if err := os.Chmod(GetXrayBinPath(), 0755); err != nil {
		return fmt.Errorf("failed to change file permissions: %w", err)
	}
	_ = filepath.Walk(constant.BinDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if ext == ".zip" || ext == ".md" || info.Name() == "LICENSE" {
				err := os.Remove(path)
				if err != nil {
					return fmt.Errorf("failed to delete file %s: %w", path, err)
				}
			}
		}
		return nil
	})
	return nil
}
