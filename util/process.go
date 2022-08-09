package util

import (
	"errors"
	"fmt"
	"runtime"
	"trojan-panel-core/module/constant"
)

func GetBinaryFile(binaryType int) (string, error) {
	binaryFile, err := GetBinaryFilePath(binaryType)
	if err != nil {
		return "", err
	}
	if !Exists(binaryFile) {
		return "", errors.New(constant.BinaryFileNotExist)
	}
	return binaryFile, nil
}

func GetBinaryFilePath(binaryType int) (string, error) {
	var binaryPath string
	var binaryName string
	switch binaryType {
	case 1:
		binaryName = fmt.Sprintf("xray-%s-%s", runtime.GOOS, runtime.GOARCH)
		binaryPath = constant.XrayPath
	case 2:
		binaryName = fmt.Sprintf("trojan-go-%s-%s", runtime.GOOS, runtime.GOARCH)
		binaryPath = constant.TrojanGoPath
	case 3:
		binaryName = fmt.Sprintf("hysteria-%s-%s", runtime.GOOS, runtime.GOARCH)
		binaryPath = constant.HysteriaPath
	default:
		return "", errors.New(constant.BinaryFileNotExist)
	}
	return fmt.Sprintf("%s/%s", binaryPath, binaryName), nil
}

func GetConfigFile(binaryType int, apiPort string) (string, error) {
	configFile, err := GetConfigFilePath(binaryType, apiPort)
	if err != nil {
		return "", err
	}
	if !Exists(configFile) {
		return "", errors.New(constant.ConfigFileNotExist)
	}
	return configFile, nil
}

func GetConfigFilePath(binaryType int, apiPort string) (string, error) {
	var configPath string
	var configName string
	switch binaryType {
	case 1:
		configName = "config.json"
		configPath = constant.XrayPath
	case 2:
		configName = fmt.Sprintf("config-%s.json", apiPort)
		configPath = constant.TrojanGoPath
	case 3:
		configName = fmt.Sprintf("config-%s.json", apiPort)
		configPath = constant.HysteriaPath
	default:
		return "", errors.New(constant.ConfigFileNotExist)
	}
	return fmt.Sprintf("%s/%s", configPath, configName), nil
}
