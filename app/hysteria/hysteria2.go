package hysteria

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"trojan-core/core"
	"trojan-core/core/process"
	"trojan-core/model/bo"
	"trojan-core/model/constant"
	"trojan-core/model/dto"
	"trojan-core/util"
)

func InitHysteria2App() error {
	apiPorts, err := util.GetConfigApiPorts(constant.Hysteria2Path)
	if err != nil {
		return err
	}
	hysteria2Process := process.NewHysteria2Instance()
	for _, apiPort := range apiPorts {
		if err != nil {
			return err
		}
		if err = hysteria2Process.StartHysteria2(apiPort); err != nil {
			return err
		}
	}
	return nil
}

func StartHysteria2(hysteria2ConfigDto dto.Hysteria2ConfigDto) error {
	var err error
	if err = initHysteria2(hysteria2ConfigDto); err != nil {
		return err
	}
	if err = process.NewHysteria2Instance().StartHysteria2(hysteria2ConfigDto.ApiPort); err != nil {
		return err
	}
	return nil
}

func StopHysteria2(apiPort uint, removeFile bool) error {
	if err := process.NewHysteria2Instance().Stop(apiPort, removeFile); err != nil {
		logrus.Errorf("hysteria stop err: %v", err)
		return err
	}
	return nil
}

func RestartHysteria2(apiPort uint) error {
	if err := StopHysteria2(apiPort, false); err != nil {
		return err
	}
	if err := StartHysteria2(dto.Hysteria2ConfigDto{ApiPort: apiPort}); err != nil {
		return err
	}
	return nil
}

func initHysteria2(hysteria2ConfigDto dto.Hysteria2ConfigDto) error {
	hysteria2ConfigFilePath, err := util.GetConfigFilePath(constant.Hysteria2, hysteria2ConfigDto.ApiPort)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(hysteria2ConfigFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logrus.Errorf("create hysteria file %s err: %v", hysteria2ConfigFilePath, err)
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
  "tls": {
    "cert": "${crt_path}",
    "key": "${key_path}"
  },
  "bandwidth": {
    "up": "${up_mbps} mbps",
    "down": "${down_mbps} mbps"
  },
  "auth": {
    "type": "http",
    "http": {
      "url": "http://127.0.0.1:${server_port}/api/auth/hysteria",
      "insecure": true
    }
  },
  "trafficStats": {
    "listen": ":${api_port}"
  }
}`
	configContent = strings.ReplaceAll(configContent, "${port}", strconv.FormatInt(int64(hysteria2ConfigDto.Port), 10))
	configContent = strings.ReplaceAll(configContent, "${crt_path}", certConfig.CrtPath)
	configContent = strings.ReplaceAll(configContent, "${key_path}", certConfig.KeyPath)
	configContent = strings.ReplaceAll(configContent, "${up_mbps}", strconv.FormatInt(int64(hysteria2ConfigDto.UpMbps), 10))
	configContent = strings.ReplaceAll(configContent, "${down_mbps}", strconv.FormatInt(int64(hysteria2ConfigDto.DownMbps), 10))
	configContent = strings.ReplaceAll(configContent, "${server_port}", strconv.FormatInt(int64(core.Config.ServerConfig.Port), 10))
	configContent = strings.ReplaceAll(configContent, "${api_port}", strconv.FormatInt(int64(hysteria2ConfigDto.ApiPort), 10))
	if len(hysteria2ConfigDto.ObfsPassword) >= 4 {
		// set obfs
		var hysteria2Config bo.Hysteria2Config
		if err = json.Unmarshal([]byte(configContent), &hysteria2Config); err != nil {
			logrus.Errorf("hysteria json.Unmarshal err: %v", err)
			return err
		}
		var hysteria2ConfigObfs bo.Hysteria2ConfigObfs
		hysteria2ConfigObfs.Type = "salamander"
		hysteria2ConfigObfs.Salamander.Password = hysteria2ConfigDto.ObfsPassword
		hysteria2Config.Obfs = &hysteria2ConfigObfs
		hysteria2ConfigStr, err := json.MarshalIndent(hysteria2Config, "", "    ")
		if err != nil {
			logrus.Errorf("hysteria json.MarshalIndent err: %v", err)
			return err
		}
		configContent = string(hysteria2ConfigStr)
	}
	_, err = file.WriteString(configContent)
	if err != nil {
		logrus.Errorf("hysteria config.json file write err: %v", err)
		return err
	}
	return nil
}

func InitHysteria2BinFile() error {
	hysteria2Path := constant.Hysteria2Path
	if !util.Exists(hysteria2Path) {
		if err := os.MkdirAll(hysteria2Path, os.ModePerm); err != nil {
			logrus.Errorf("create hysteria folder err: %v", err)
			return err
		}
	}

	binaryFilePath, err := util.GetBinaryFilePath(constant.Hysteria2)
	if err != nil {
		return err
	}
	if !util.Exists(binaryFilePath) {
		logrus.Errorf("hysteria binary does not exist")
		return errors.New(constant.BinaryFileNotExist)
	}
	return nil
}
