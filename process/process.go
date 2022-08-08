package process

import (
	"errors"
	"fmt"
	"runtime"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

func GetBinaryFilePath(name string) (string, error) {
	var binaryName string
	var binaryPath string
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
		return "", errors.New(constant.BinaryNotExist)
	}
	binaryFilePath := fmt.Sprintf("%s/%s", binaryPath, binaryName)
	if !util.Exists(binaryFilePath) {
		return "", errors.New(constant.BinaryNotExist)
	}
	return binaryFilePath, nil
}
