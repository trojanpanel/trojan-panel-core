package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"trojan-panel-core/module/constant"
)

func GetXrayConfigFileNameByApiPort(apiPort uint, protocol string) (string, error) {
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
