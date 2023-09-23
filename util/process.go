package util

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"trojan-panel-core/model/constant"
)

var configFileNameReg = regexp.MustCompile("^config-([1-9]\\d*)[\\s\\S]*\\.json$")

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
	case constant.Xray:
		binaryName = "xray"
		binaryPath = constant.XrayBinPath
	case constant.TrojanGo:
		binaryName = "trojan-go"
		binaryPath = constant.TrojanGoBinPath
	case constant.Hysteria:
		binaryName = "hysteria"
		binaryPath = constant.HysteriaBinPath
	case constant.NaiveProxy:
		binaryName = "naiveproxy"
		binaryPath = constant.NaiveProxyBinPath
	case constant.Hysteria2:
		binaryName = "hysteria2"
		binaryPath = constant.Hysteria2BinPath
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
	case constant.Xray:
		configPath = constant.XrayPath
		var err error
		configFileName, err = GetXrayConfigFileNameByApiPort(apiPort)
		if err != nil {
			return "", err
		}
	case constant.TrojanGo:
		configPath = constant.TrojanGoPath
		configFileName = fmt.Sprintf("config-%d.json", apiPort)
	case constant.Hysteria:
		configPath = constant.HysteriaPath
		configFileName = fmt.Sprintf("config-%d.json", apiPort)
	case constant.NaiveProxy:
		configPath = constant.NaiveProxyPath
		configFileName = fmt.Sprintf("config-%d.json", apiPort)
	case constant.Hysteria2:
		configPath = constant.Hysteria2Path
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

func GetConfigApiPorts(dirPth string) ([]uint, error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	apiPorts := make([]uint, 0)
	for _, fi := range dir {
		// filter specified format
		finds := configFileNameReg.FindStringSubmatch(fi.Name())
		if len(finds) > 0 {
			apiPort, err := strconv.Atoi(finds[1])
			if err != nil {
				logrus.Errorf("type conversion err: %v", err)
				continue
			}
			apiPorts = append(apiPorts, uint(apiPort))
		}
	}
	return apiPorts, nil
}
