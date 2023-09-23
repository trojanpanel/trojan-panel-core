package hysteria

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"trojan-panel-core/core"
	"trojan-panel-core/core/process"
	"trojan-panel-core/model/constant"
	"trojan-panel-core/model/dto"
	"trojan-panel-core/util"
)

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

func StopHysteria(apiPort uint, removeFile bool) error {
	if err := process.NewHysteriaInstance().Stop(apiPort, removeFile); err != nil {
		logrus.Errorf("hysteria stop err: %v", err)
		return err
	}
	return nil
}

func RestartHysteria(apiPort uint) error {
	if err := StopHysteria(apiPort, false); err != nil {
		return err
	}
	if err := StartHysteria(dto.HysteriaConfigDto{ApiPort: apiPort}); err != nil {
		return err
	}
	return nil
}

func initHysteria(hysteriaConfigDto dto.HysteriaConfigDto) error {
	hysteriaConfigFilePath, err := util.GetConfigFilePath(constant.Hysteria, hysteriaConfigDto.ApiPort)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(hysteriaConfigFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logrus.Errorf("create hysteria file %s err: %v", hysteriaConfigFilePath, err)
		return err
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	certConfig := core.Config.CertConfig
	configContent := `{
  "listen": ":${port}",
  "protocol": "${protocol}",
  "cert": "${crt_path}",
  "key": "${key_path}",
  "obfs": "${obfs}",
  "up_mbps": ${up_mbps},
  "down_mbps": ${down_mbps},
  "auth": {
    "mode": "external",
    "config": {
      "http": "http://127.0.0.1:${server_port}/api/auth/hysteria"
    }
  }
}`
	configContent = strings.ReplaceAll(configContent, "${port}", strconv.FormatInt(int64(hysteriaConfigDto.Port), 10))
	configContent = strings.ReplaceAll(configContent, "${protocol}", hysteriaConfigDto.Protocol)
	configContent = strings.ReplaceAll(configContent, "${crt_path}", certConfig.CrtPath)
	configContent = strings.ReplaceAll(configContent, "${key_path}", certConfig.KeyPath)
	configContent = strings.ReplaceAll(configContent, "${obfs}", hysteriaConfigDto.Obfs)
	configContent = strings.ReplaceAll(configContent, "${up_mbps}", strconv.FormatInt(int64(hysteriaConfigDto.UpMbps), 10))
	configContent = strings.ReplaceAll(configContent, "${down_mbps}", strconv.FormatInt(int64(hysteriaConfigDto.DownMbps), 10))
	configContent = strings.ReplaceAll(configContent, "${server_port}", strconv.FormatInt(int64(core.Config.ServerConfig.Port), 10))
	_, err = file.WriteString(configContent)
	if err != nil {
		logrus.Errorf("hysteria config.json file write err: %v", err)
		return err
	}
	return nil
}

func InitHysteriaBinFile() error {
	hysteriaPath := constant.HysteriaPath
	if !util.Exists(hysteriaPath) {
		if err := os.MkdirAll(hysteriaPath, os.ModePerm); err != nil {
			logrus.Errorf("create hysteria folder err: %v", err)
			return err
		}
	}

	binaryFilePath, err := util.GetBinaryFilePath(constant.Hysteria)
	if err != nil {
		return err
	}
	if !util.Exists(binaryFilePath) {
		logrus.Errorf("hysteria binary does not exist")
		return errors.New(constant.BinaryFileNotExist)
	}
	return nil
}
