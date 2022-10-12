package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
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

func GetConfigFile(binaryType int, apiPort uint) (string, error) {
	configFile, err := GetConfigFilePath(binaryType, apiPort)
	if err != nil {
		return "", err
	}
	if !Exists(configFile) {
		return "", errors.New(constant.ConfigFileNotExist)
	}
	return configFile, nil
}

func GetConfigFilePath(binaryType int, apiPort uint) (string, error) {
	var configPath string
	var configFileName string
	switch binaryType {
	case 1:
		configPath = constant.XrayPath
		var err error
		configFileName, err = GetXrayConfigFileNameByApiPort(apiPort)
		if err != nil {
			return "", err
		}
	case 2:
		configPath = constant.TrojanGoPath
		configFileName = fmt.Sprintf("config-%d.json", apiPort)
	case 3:
		configPath = constant.HysteriaPath
		configFileName = fmt.Sprintf("config-%d.json", apiPort)
	default:
		return "", errors.New(constant.ConfigFileNotExist)
	}
	return fmt.Sprintf("%s/%s", configPath, configFileName), nil
}

func GetXrayConfigFileNameByApiPort(apiPort uint) (string, error) {
	fileNamePrefix := fmt.Sprintf("config-%d", apiPort)
	dir, err := ioutil.ReadDir(constant.XrayPath)
	if err != nil {
		return "", err
	}
	for _, fi := range dir {
		if strings.HasPrefix(fi.Name(), fileNamePrefix) {
			return fi.Name(), nil
		}
	}
	return "", errors.New(constant.ConfigFileNotExist)
}

func GetXrayProtocolByApiPort(apiPort uint) (string, error) {
	xrayConfigFileName, err := GetXrayConfigFileNameByApiPort(apiPort)
	if err != nil {
		return "", err
	}
	start := strings.LastIndex(xrayConfigFileName, "-") + 1
	end := strings.LastIndex(xrayConfigFileName, ".")
	return xrayConfigFileName[start:end], nil
}
