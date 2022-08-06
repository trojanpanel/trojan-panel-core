package trojango

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"sync"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/pkg/trojango/start"
	"trojan-panel-core/util"
)

var trojanGoMap sync.Map

func StartTrojanGo() {
	go func() {
		start.TrojanGoMain()
	}()
}

func StopTrojanGo() {

}

func ConfigTrojanGo(trojanGoConfigDto dto.TrojanGoConfigDto) {
	trojanGoConfigFilePath := constant.TrojanGoConfigFilePath
	if !util.Exists(trojanGoConfigFilePath) {
		file, err := os.Create(trojanGoConfigFilePath)
		if err != nil {
			logrus.Errorf("创建trojan go config.json文件异常 err: %v\n", err)
			panic(err)
		}
		defer file.Close()

		configContent := `{
    "run_type": "server",
    "local_addr": "0.0.0.0",
    "local_port": ${local_port},
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
        "cert": "/tpdata/caddy/acme/${domain}/${domain}.crt",
        "key": "/tpdata/caddy/acme/${domain}/${domain}.key",
        "key_password": "",
        "cipher": "",
        "curves": "",
        "prefer_server_cipher": false,
        "sni": "",
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
        "host": "${domain}"
    },
    "shadowsocks": {
        "enabled": ${shadowsocks_enabled},
        "method": "${shadowsocks_method}",
        "password": "${shadowsocks_password}"
    },
    "api": {
        "enabled": true,
        "api_addr": "127.0.0.1",
        "api_port": ${api_port}
    }
}
`
		configContent = strings.ReplaceAll(configContent, "${local_port}", trojanGoConfigDto.LocalPort)
		configContent = strings.ReplaceAll(configContent, "${domain}", trojanGoConfigDto.Domain)
		configContent = strings.ReplaceAll(configContent, "${mux_enable}", trojanGoConfigDto.MuxEnable)
		configContent = strings.ReplaceAll(configContent, "${websocket_enabled}", trojanGoConfigDto.WebsocketEnable)
		configContent = strings.ReplaceAll(configContent, "${websocket_path}", trojanGoConfigDto.WebsocketPath)
		configContent = strings.ReplaceAll(configContent, "${shadowsocks_enabled}", trojanGoConfigDto.ShadowsocksEnable)
		configContent = strings.ReplaceAll(configContent, "${shadowsocks_method}", trojanGoConfigDto.ShadowsocksMethod)
		configContent = strings.ReplaceAll(configContent, "${shadowsocks_password}", trojanGoConfigDto.ShadowsocksPassword)
		configContent = strings.ReplaceAll(configContent, "${api_port}", trojanGoConfigDto.ApiPort)
		_, err = file.WriteString(configContent)
		if err != nil {
			logrus.Errorf("trojan go config.json文件写入异常 err: %v\n", err)
			panic(err)
		}
	}
}
