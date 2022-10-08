package trojango

import (
	"errors"
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

// InitTrojanGoApp 初始化TrojanGo应用
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

// StartTrojanGo 启动TrojanGo
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

// StopTrojanGo 暂停TrojanGo
func StopTrojanGo(apiPort uint, removeFile bool) error {
	if err := process.NewTrojanGoInstance().Stop(apiPort, removeFile); err != nil {
		logrus.Errorf("trojan go stop err: %v\n", err)
		return err
	}
	return nil
}

// RestartTrojanGo 重启TrojanGo
func RestartTrojanGo(apiPort uint) error {
	if err := StopTrojanGo(apiPort, false); err != nil {
		return err
	}
	if err := StartTrojanGo(dto.TrojanGoConfigDto{ApiPort: apiPort}); err != nil {
		return err
	}
	return nil
}

// 初始化TrojanGo文件
func initTrojanGo(trojanGoConfigDto dto.TrojanGoConfigDto) error {
	// 初始化文件夹
	trojanGoPath := constant.TrojanGoPath
	if !util.Exists(trojanGoPath) {
		if err := os.MkdirAll(trojanGoPath, os.ModePerm); err != nil {
			logrus.Errorf("创建Trojan Go文件夹异常 err: %v\n", err)
			panic(err)
		}
	}

	// 下载二进制文件
	binaryFilePath, err := util.GetBinaryFilePath(2)
	if err != nil {
		return err
	}
	if !util.Exists(binaryFilePath) {
		if err = util.DownloadFile(fmt.Sprintf("%s/trojan-go-%s-%s", constant.DownloadBaseUrl, runtime.GOOS, runtime.GOARCH),
			binaryFilePath); err != nil {
			logrus.Errorf("Trojan Go二进制文件下载失败 err: %v\n", err)
			panic(errors.New(constant.DownloadFilError))
		}
	}

	// 初始化配置
	trojanGoConfigFilePath, err := util.GetConfigFilePath(2, trojanGoConfigDto.ApiPort)
	if err != nil {
		return err
	}
	if !util.Exists(trojanGoConfigFilePath) {
		file, err := os.Create(trojanGoConfigFilePath)
		if err != nil {
			logrus.Errorf("创建trojan go config-%d.json文件异常 err: %v\n", trojanGoConfigDto.ApiPort, err)
			panic(err)
		}
		defer file.Close()

		certConfig := core.Config.CertConfig

		configContent := `{
  "run_type": "server",
  "local_addr": "0.0.0.0",
  "local_port": ${port},
  "remote_addr": "trojan-panel-caddy",
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
    "fallback_addr": "",
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
			logrus.Errorf("trojan go config-%d.json文件写入异常 err: %v\n", trojanGoConfigDto.ApiPort, err)
			panic(err)
		}
	}
	return nil
}
