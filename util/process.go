package util

import (
	"errors"
	"fmt"
	"runtime"
	"trojan-panel-core/module/constant"
)

func GetBinaryFile(name string) (string, error) {
	binaryFile, err := GetBinaryFilePath(name)
	if err != nil {
		return "", err
	}
	if !Exists(binaryFile) {
		return "", errors.New(constant.BinaryFileNotExist)
	}
	return binaryFile, nil
}

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
	return fmt.Sprintf("%s/%s", binaryPath, binaryName), nil
}

func GetConfigFile(apiPort string, name string) (string, error) {
	configFile, err := GetConfigFilePath(apiPort, name)
	if err != nil {
		return "", err
	}
	if !Exists(configFile) {
		return "", errors.New(constant.ConfigFileNotExist)
	}
	return configFile, nil
}

func GetConfigFilePath(apiPort string, name string) (string, error) {
	var configPath string
	var configName string
	switch name {
	case "xray":
		configName = "config.json"
		configPath = constant.XrayPath
	case "trojan-go":
		configName = fmt.Sprintf("config-%s.json", apiPort)
		configPath = constant.TrojanGoPath
	case "hysteria":
		configName = fmt.Sprintf("config-%s.json", apiPort)
		configPath = constant.HysteriaPath
	default:
		return "", errors.New(constant.ConfigFileNotExist)
	}
	return fmt.Sprintf("%s/%s", configPath, configName), nil
}
