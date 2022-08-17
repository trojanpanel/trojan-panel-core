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
	"trojan-panel-core/dao"
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
		if err = syncTrojanGoData(apiPort); err != nil {
			return err
		}
	}
	return nil
}

// 数据库同步至应用
func syncTrojanGoData(apiPort uint) error {
	api := NewTrojanGoApi(apiPort)
	apiUserVos, err := dao.SelectUsersToApi(true)
	if err != nil {
		return err
	}
	for _, apiUser := range apiUserVos {
		userDto := dto.TrojanGoAddUserDto{
			Password:        apiUser.Password,
			DownloadTraffic: apiUser.Download,
			UploadTraffic:   apiUser.Upload,
		}
		if err := api.AddUser(userDto); err != nil {
			logrus.Errorf("数据库同步至应用 trojan go api用户添加失败 err:%v\n", err)
			continue
		}
	}
	apiUserVos, err = dao.SelectUsersToApi(false)
	if err != nil {
		return err
	}
	for _, apiUser := range apiUserVos {
		if err := api.DeleteUser(apiUser.Password); err != nil {
			logrus.Errorf("数据库同步至应用 trojan go api用户删除失败 err:%v\n", err)
			continue
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
func StopTrojanGo(apiPort uint) error {
	if err := process.NewTrojanGoInstance().Stop(apiPort); err != nil {
		logrus.Errorf("trojan go stop err: %v\n", err)
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
    "enabled": ${websocket_enable},
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
	"api_addr": "",
	"api_port": ${apiPort}
  }
}
`
		configContent = strings.ReplaceAll(configContent, "${port}", strconv.FormatInt(int64(trojanGoConfigDto.Port), 10))
		configContent = strings.ReplaceAll(configContent, "${ip}", trojanGoConfigDto.Ip)
		configContent = strings.ReplaceAll(configContent, "${sni}", trojanGoConfigDto.Sni)
		configContent = strings.ReplaceAll(configContent, "${mux_enable}", trojanGoConfigDto.MuxEnable)
		configContent = strings.ReplaceAll(configContent, "${websocket_enable}", trojanGoConfigDto.WebsocketEnable)
		configContent = strings.ReplaceAll(configContent, "${websocket_path}", trojanGoConfigDto.WebsocketPath)
		configContent = strings.ReplaceAll(configContent, "${websocket_host}", trojanGoConfigDto.WebsocketHost)
		configContent = strings.ReplaceAll(configContent, "${ss_enable}", trojanGoConfigDto.SSEnable)
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
