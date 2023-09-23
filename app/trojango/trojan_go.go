package trojango

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

func InitTrojanGoApp() error {
	apiPorts, err := util.GetConfigApiPorts(constant.TrojanGoPath)
	if err != nil {
		return err
	}
	trojanGoInstance := process.NewTrojanGoInstance()
	for _, apiPort := range apiPorts {
		if err = trojanGoInstance.StartTrojanGo(apiPort); err != nil {
			return err
		}
	}
	return nil
}

func StartTrojanGo(trojanGoConfigDto dto.TrojanGoConfigDto) error {
	var err error
	if err = initTrojanGo(trojanGoConfigDto); err != nil {
		return err
	}
	if err = process.NewTrojanGoInstance().StartTrojanGo(trojanGoConfigDto.ApiPort); err != nil {
		return err
	}
	return nil
}

func StopTrojanGo(apiPort uint, removeFile bool) error {
	if err := process.NewTrojanGoInstance().Stop(apiPort, removeFile); err != nil {
		logrus.Errorf("trojango stop err: %v", err)
		return err
	}
	return nil
}

func RestartTrojanGo(apiPort uint) error {
	if err := StopTrojanGo(apiPort, false); err != nil {
		return err
	}
	if err := StartTrojanGo(dto.TrojanGoConfigDto{ApiPort: apiPort}); err != nil {
		return err
	}
	return nil
}

func initTrojanGo(trojanGoConfigDto dto.TrojanGoConfigDto) error {
	trojanGoConfigFilePath, err := util.GetConfigFilePath(constant.TrojanGo, trojanGoConfigDto.ApiPort)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(trojanGoConfigFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		logrus.Errorf("create trojango file %s err: %v", trojanGoConfigFilePath, err)
		return err
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	certConfig := core.Config.CertConfig

	configContent := `{
  "run_type": "server",
  "local_addr": "0.0.0.0",
  "local_port": ${port},
  "remote_addr": "${remote_addr}",
  "remote_port": 80,
  "log_level": 1,
  "log_file": "",
  "password": [],
  "disable_http_check": false,
  "udp_timeout": 60,
  "ssl": {
    "verify": true,
    "verify_hostname": true,
    "cert": "${crt_path}",
    "key": "${key_path}",
    "key_password": "",
    "cipher": "",
    "curves": "",
    "prefer_server_cipher": false,
    "sni": "${sni}",
    "alpn": [
      "http/1.1"
    ],
    "session_ticket": true,
    "reuse_session": true,
    "plain_http_response": "",
    "fallback_addr": "${fallback_addr}",
    "fallback_port": 80,
    "fingerprint": ""
  },
  "tcp": {
    "no_delay": true,
    "keep_alive": true,
    "prefer_ipv4": false
  },
    "mux": {
    "enabled": ${mux_enable},
    "concurrency": 8,
    "idle_timeout": 60
  },
  "websocket": {
    "enabled": ${websocket_enable},
    "path": "${websocket_path}",
    "host": "${websocket_host}"
  },
  "shadowsocks": {
    "enabled": ${ss_enable},
    "method": "${ss_method}",
    "password": "${ss_password}"
  },
  "api": {
	"enabled": true,
	"api_addr": "127.0.0.1",
	"api_port": ${api_port},
	"ssl": {
      "enabled": false,
      "key": "",
      "cert": "",
      "verify_client": false,
      "client_cert": []
    }
  }
}
`
	configContent = strings.ReplaceAll(configContent, "${port}", strconv.FormatInt(int64(trojanGoConfigDto.Port), 10))
	configContent = strings.ReplaceAll(configContent, "${crt_path}", certConfig.CrtPath)
	configContent = strings.ReplaceAll(configContent, "${key_path}", certConfig.KeyPath)
	configContent = strings.ReplaceAll(configContent, "${sni}", trojanGoConfigDto.Sni)
	// custom cert points to Microsoft
	if certConfig.CrtPath != "" && strings.Contains(certConfig.CrtPath, "custom_cert") {
		configContent = strings.ReplaceAll(configContent, "${remote_addr}", "www.microsoft.com")
		configContent = strings.ReplaceAll(configContent, "${fallback_addr}", "www.microsoft.com")
	} else {
		configContent = strings.ReplaceAll(configContent, "${remote_addr}", "127.0.0.1")
		configContent = strings.ReplaceAll(configContent, "${fallback_addr}", "127.0.0.1")
	}
	var muxEnableStr string
	if trojanGoConfigDto.MuxEnable == 1 {
		muxEnableStr = "true"
	} else {
		muxEnableStr = "false"
	}
	configContent = strings.ReplaceAll(configContent, "${mux_enable}", muxEnableStr)
	var websocketEnableStr string
	if trojanGoConfigDto.WebsocketEnable == 1 {
		websocketEnableStr = "true"
	} else {
		websocketEnableStr = "false"
	}
	configContent = strings.ReplaceAll(configContent, "${websocket_enable}", websocketEnableStr)
	configContent = strings.ReplaceAll(configContent, "${websocket_path}", trojanGoConfigDto.WebsocketPath)
	configContent = strings.ReplaceAll(configContent, "${websocket_host}", trojanGoConfigDto.WebsocketHost)
	var ssEnableStr string
	if trojanGoConfigDto.SSEnable == 1 {
		ssEnableStr = "true"
	} else {
		ssEnableStr = "false"
	}
	configContent = strings.ReplaceAll(configContent, "${ss_enable}", ssEnableStr)
	configContent = strings.ReplaceAll(configContent, "${ss_method}", trojanGoConfigDto.SSMethod)
	configContent = strings.ReplaceAll(configContent, "${ss_password}", trojanGoConfigDto.SSPassword)
	configContent = strings.ReplaceAll(configContent, "${api_port}", strconv.FormatInt(int64(trojanGoConfigDto.ApiPort), 10))
	_, err = file.WriteString(configContent)
	if err != nil {
		logrus.Errorf("trojango file config-%d.json write err: %v", trojanGoConfigDto.ApiPort, err)
		return err
	}
	return nil
}

func InitTrojanGoBinFile() error {
	trojanGoPath := constant.TrojanGoPath
	if !util.Exists(trojanGoPath) {
		if err := os.MkdirAll(trojanGoPath, os.ModePerm); err != nil {
			logrus.Errorf("create trojango folder err: %v", err)
			return err
		}
	}

	binaryFilePath, err := util.GetBinaryFilePath(constant.TrojanGo)
	if err != nil {
		return err
	}
	if !util.Exists(binaryFilePath) {
		logrus.Errorf("trojango binary file does not exist")
		return errors.New(constant.BinaryFileNotExist)
	}
	return nil
}
