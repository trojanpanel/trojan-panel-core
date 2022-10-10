package hysteria

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
	"trojan-panel-core/core"
	"trojan-panel-core/core/process"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/util"
)

// InitHysteriaApp 初始化Hysteria应用
func InitHysteriaApp() error {
	apiPorts, err := util.GetConfigApiPorts(constant.HysteriaPath)
	if err != nil {
		return err
	}
	hysteriaProcess := process.NewHysteriaInstance()
	for _, apiPort := range apiPorts {
		if err != nil {
			return err
		}
		if err = hysteriaProcess.StartHysteria(apiPort); err != nil {
			return err
		}
	}
	return nil
}

// StartHysteria 启动Hysteria
func StartHysteria(hysteriaConfigDto dto.HysteriaConfigDto) error {
	var err error
	if err = initHysteria(hysteriaConfigDto); err != nil {
		return err
	}
	if err = process.NewHysteriaInstance().StartHysteria(hysteriaConfigDto.ApiPort); err != nil {
		return err
	}
	return nil
}

// StopHysteria 暂停Hysteria
func StopHysteria(apiPort uint, removeFile bool) error {
	if err := process.NewHysteriaInstance().Stop(apiPort, removeFile); err != nil {
		logrus.Errorf("hysteria stop err: %v\n", err)
		return err
	}
	return nil
}

// RestartHysteria 重启Hysteria
func RestartHysteria(apiPort uint) error {
	if err := StopHysteria(apiPort, false); err != nil {
		return err
	}
	if err := StartHysteria(dto.HysteriaConfigDto{ApiPort: apiPort}); err != nil {
		return err
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
			return err
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
			return err
		}
	}

	// 初始化配置
	hysteriaConfigFilePath, err := util.GetConfigFilePath(3, hysteriaConfigDto.ApiPort)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(hysteriaConfigFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logrus.Errorf("创建hysteria %s文件异常 err: %v\n", hysteriaConfigFilePath, err)
		return err
	}
	defer file.Close()

	certConfig := core.Config.CertConfig
	configContent := `{
  "listen": ":${port}",
  "protocol": "${protocol}",
  "cert": "${crt_path}",
  "key": "${key_path}",
  "up_mbps": ${up_mbps},
  "down_mbps": ${down_mbps},
  "auth": {
    "mode": "external",
    "config": {
      "http": "http://127.0.0.1:8082/api/auth/hysteria"
    }
  }
}`
	configContent = strings.ReplaceAll(configContent, "${port}", strconv.FormatInt(int64(hysteriaConfigDto.Port), 10))
	configContent = strings.ReplaceAll(configContent, "${protocol}", hysteriaConfigDto.Protocol)
	configContent = strings.ReplaceAll(configContent, "${crt_path}", certConfig.CrtPath)
	configContent = strings.ReplaceAll(configContent, "${key_path}", certConfig.KeyPath)
	configContent = strings.ReplaceAll(configContent, "${up_mbps}", strconv.FormatInt(int64(hysteriaConfigDto.UpMbps), 10))
	configContent = strings.ReplaceAll(configContent, "${down_mbps}", strconv.FormatInt(int64(hysteriaConfigDto.DownMbps), 10))
	_, err = file.WriteString(configContent)
	if err != nil {
		logrus.Errorf("hysteria config.json文件写入异常 err: %v\n", err)
		return err
	}
	return nil
}
