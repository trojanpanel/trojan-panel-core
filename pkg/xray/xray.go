package xray

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

// StartXray 启动Xray
func StartXray() error {
	if err := InitXray(constant.GrpcPortXray); err != nil {
		return errors.New(constant.SysError)
	}
	return nil
}

// StopXray 关闭Xray
func StopXray() {

}

func InitXray(apiPort string) error {
	// 初始化文件夹
	xrayPath := constant.XrayPath
	if !util.Exists(xrayPath) {
		if err := os.MkdirAll(xrayPath, os.ModePerm); err != nil {
			logrus.Errorf("创建Xray文件夹异常 err: %v\n", err)
			return err
		}
	}
	// 初始化配置
	xrayConfigFilePath := constant.XrayConfigFilePath
	if !util.Exists(xrayConfigFilePath) {
		file, err := os.Create(xrayConfigFilePath)
		if err != nil {
			logrus.Errorf("创建xray config.json文件异常 err: %v\n", err)
			panic(err)
		}
		defer file.Close()

		configContent := `{
  "stats": {},
  "api": {
    "services": [
      "HandlerService",
      "LoggerService",
      "StatsService"
    ],
    "tag": "api"
  },
  "policy": {
    "levels": {
      "0": {
        "statsUserUplink": true,
        "statsUserDownlink": true
      }
    },
    "system": {
      "statsInboundUplink": true,
      "statsInboundDownlink": true,
      "statsOutboundUplink": true,
      "statsOutboundDownlink": true
    }
  },
  "inbounds": [
    {
      "tag": "api",
      "listen": "127.0.0.1",
      "port": ${api_port},
      "protocol": "dokodemo-door",
      "settings": {
        "address": "127.0.0.1"
      }
    }
  ],
  "outbounds": [
    {
      "tag": "direct",
      "protocol": "freedom",
      "settings": {}
    }
  ],
  "routing": {
    "rules": [
      {
        "inboundTag": [
          "api"
        ],
        "outboundTag": "api",
        "type": "field"
      }
    ]
  }
}
`
		configContent = strings.ReplaceAll(configContent, "${api_port}", apiPort)
		_, err = file.WriteString(configContent)
		if err != nil {
			logrus.Errorf("xray config.json文件写入异常 err: %v\n", err)
			return err
		}
	}
	return nil
}
