package util

import (
	"errors"
	"fmt"
	"runtime"
	"trojan-panel-core/module/constant"
)

func GetBinaryFilePath(name string) (string, error) {
	var binaryPath string
	var binaryName string
	switch name {
	case "xray":
		binaryName = fmt.Sprintf("xray-%s-%s", runtime.GOOS, runtime.GOARCH)
		binaryPath = constant.XrayPath
	case "trojan-go":
		binaryName = fmt.Sprintf("trojan-go-%s-%s", runtime.GOOS, runtime.GOARCH)
		binaryPath = constant.TrojanGoPath
	case "hysteria":
		binaryName = fmt.Sprintf("hysteria-%s-%s", runtime.GOOS, runtime.GOARCH)
		binaryPath = constant.HysteriaPath
	default:
		return "", errors.New(constant.BinaryFileNotExist)
	}
	binaryFilePath := fmt.Sprintf("%s/%s", binaryPath, binaryName)
	if !Exists(binaryFilePath) {
		return "", errors.New(constant.BinaryFileNotExist)
	}
	return binaryFilePath, nil
}

func GetConfigPath(id int, name string) (string, error) {
	var configPath string
	var configName string
	switch name {
	case "xray":
		configName = "config.json"
		configPath = constant.XrayPath
	case "trojan-go":
		configName = fmt.Sprintf("trojan-go-config-%d.json", id)
		configPath = constant.TrojanGoPath
	case "hysteria":
		configName = fmt.Sprintf("hysteria-config-%d.json", id)
		configPath = constant.HysteriaPath
	default:
		return "", errors.New(constant.ConfigFileNotExist)
	}
	configFilePath := fmt.Sprintf("%s/%s", configPath, configName)
	if !Exists(configFilePath) {
		return "", errors.New(constant.ConfigFileNotExist)
	}
	return configFilePath, nil
}
