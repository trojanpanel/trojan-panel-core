package util

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strconv"
	"trojan-core/model/constant"
)

var configFileNameReg = regexp.MustCompile("^([1-9]\\d*)\\.json$")

func GetApiPorts(dirPth string) ([]uint, error) {
	dir, err := os.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	apiPorts := make([]uint, 0)
	for _, fi := range dir {
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

func GetBinPath(proxy string) (string, error) {
	var binaryPath string
	var binaryName string
	switch proxy {
	case constant.Xray:
		binaryName = "xray"
		binaryPath = constant.XrayBinPath
	case constant.Hysteria:
		binaryName = "hysteria"
		binaryPath = constant.HysteriaBinPath
	case constant.Naive:
		binaryName = "naive"
		binaryPath = constant.NaiveBinPath
	default:
		return "", errors.New("bin file does not exist")
	}
	return fmt.Sprintf("%s%s", binaryPath, binaryName), nil
}

func GetConfigPath(proxy string, apiPort uint) (string, error) {
	var configPath string
	switch proxy {
	case constant.Xray:
		configPath = constant.XrayPath
	case constant.Hysteria:
		configPath = constant.Hysteria
	case constant.Naive:
		configPath = constant.Naive
	default:
		return "", errors.New("config file does not exist")
	}
	return fmt.Sprintf("%s%d.json", configPath, apiPort), nil
}
