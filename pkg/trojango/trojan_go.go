package trojango

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
	"trojan-panel-core/core/process"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/util"
)

var trojanGoProcess *process.TrojanGoProcess

func StartTrojanGo(trojanGoConfigDto dto.TrojanGoConfigDto) error {
	var err error
	if err = initTrojanGo(trojanGoConfigDto); err != nil {
		return err
	}
	trojanGoProcess, err = process.NewTrojanGoProcess(trojanGoConfigDto.ApiPort)
	if err != nil {
		return err
	}
	if err = trojanGoProcess.StartTrojanGo(trojanGoConfigDto.ApiPort); err != nil {
		return err
	}
	return nil
}

func StopTrojanGo(apiPort int) error {
	if trojanGoProcess != nil {
		if err := trojanGoProcess.Stop(apiPort); err != nil {
			return err
		}
	}
	return nil
}

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
			logrus.Errorf("创建trojan go config-%d.json文件异常 err: %v\n", trojanGoConfigDto.Id, err)
			panic(err)
		}
		defer file.Close()

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
        "cert": "/tpdata/caddy/acme/${ip}/${ip}.crt",
        "key": "/tpdata/caddy/acme/${ip}/${ip}.key",
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
        "enabled": ${websocket_enabled},
        "path": "/${websocket_path}",
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
        "api_port": ${api_port}
    }
}
`
		configContent = strings.ReplaceAll(configContent, "${ip}", trojanGoConfigDto.Ip)
		configContent = strings.ReplaceAll(configContent, "${port}", trojanGoConfigDto.Port)
		configContent = strings.ReplaceAll(configContent, "${api_port}", strconv.Itoa(trojanGoConfigDto.ApiPort))
		configContent = strings.ReplaceAll(configContent, "${sni}", trojanGoConfigDto.Sni)
		configContent = strings.ReplaceAll(configContent, "${mux_enable}", trojanGoConfigDto.MuxEnable)
		configContent = strings.ReplaceAll(configContent, "${websocket_enabled}", trojanGoConfigDto.WebsocketEnable)
		configContent = strings.ReplaceAll(configContent, "${websocket_path}", trojanGoConfigDto.WebsocketPath)
		configContent = strings.ReplaceAll(configContent, "${websocket_host}", trojanGoConfigDto.WebsocketHost)
		configContent = strings.ReplaceAll(configContent, "${ss_enable}", trojanGoConfigDto.SSEnable)
		configContent = strings.ReplaceAll(configContent, "${ss_method}", trojanGoConfigDto.SSMethod)
		configContent = strings.ReplaceAll(configContent, "${ss_password}", trojanGoConfigDto.SSPassword)
		_, err = file.WriteString(configContent)
		if err != nil {
			logrus.Errorf("trojan go config-%d.json文件写入异常 err: %v\n", trojanGoConfigDto.Id, err)
			panic(err)
		}
	}
	return nil
}
