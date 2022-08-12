package hysteria

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"trojan-panel-core/core/process"
	"trojan-panel-core/dao"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/util"
)

var hysteriaProcess *process.HysteriaProcess

// InitHysteriaApp 初始化Hysteria应用
func InitHysteriaApp() error {
	apiPorts, err := util.GetConfigApiPorts(constant.HysteriaPath)
	if err != nil {
		return err
	}
	for _, apiPort := range apiPorts {
		hysteriaProcess, err := process.NewHysteriaProcess(apiPort)
		if err != nil {
			return err
		}
		if err = hysteriaProcess.StartHysteria(apiPort); err != nil {
			return err
		}
		if err = syncHysteriaData(apiPort); err != nil {
			return err
		}
	}
	return nil
}

// 数据库同步至应用
func syncHysteriaData(apiPort int) error {
	_, err := dao.SelectUsersToApi(true)
	if err != nil {
		return err
	}
	return nil
}

// StartHysteria 启动Hysteria
func StartHysteria(hysteriaConfigDto dto.HysteriaConfigDto) error {
	var err error
	if err = initHysteria(hysteriaConfigDto); err != nil {
		return err
	}
	hysteriaProcess, err = process.NewHysteriaProcess(hysteriaConfigDto.ApiPort)
	if err != nil {
		return err
	}
	if err = hysteriaProcess.StartHysteria(hysteriaConfigDto.ApiPort); err != nil {
		return err
	}
	return nil
}

// StopHysteria 暂停Hysteria
func StopHysteria(apiPort int) error {
	if hysteriaProcess != nil {
		if err := hysteriaProcess.Stop(apiPort); err != nil {
			return err
		}
	}
	return nil
}

// 初始化Hysteria文件
func initHysteria(hysteriaConfigDto dto.HysteriaConfigDto) error {
	// 初始化文件夹
	hysteriaPath := constant.HysteriaPath
	if !util.Exists(hysteriaPath) {
		if err := os.MkdirAll(hysteriaPath, os.ModePerm); err != nil {
			logrus.Errorf("创建Hysteria文件夹异常 err: %v\n", err)
			panic(err)
		}
	}

	// 下载二进制文件
	binaryFilePath, err := util.GetBinaryFilePath(3)
	if err != nil {
		return err
	}
	if !util.Exists(binaryFilePath) {
		if err = util.DownloadFile(fmt.Sprintf("%s/hysteria-%s-%s", constant.DownloadBaseUrl, runtime.GOOS, runtime.GOARCH),
			binaryFilePath); err != nil {
			logrus.Errorf("Hysteria二进制文件下载失败 err: %v\n", err)
			panic(errors.New(constant.DownloadFilError))
		}
	}

	// 初始化配置
	hysteriaConfigFilePath, err := util.GetConfigFilePath(3, hysteriaConfigDto.ApiPort)
	if err != nil {
		return err
	}
	if !util.Exists(hysteriaConfigFilePath) {
		file, err := os.Create(hysteriaConfigFilePath)
		if err != nil {
			logrus.Errorf("创建hysteria config.json文件异常 err: %v\n", err)
			panic(err)
		}
		defer file.Close()

		configContent := ``
		_, err = file.WriteString(configContent)
		if err != nil {
			logrus.Errorf("hysteria config.json文件写入异常 err: %v\n", err)
			panic(err)
		}
	}
	return nil
}
