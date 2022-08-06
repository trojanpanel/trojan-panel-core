package xray

import (
	"github.com/sirupsen/logrus"
	"os"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

func StartXray() error {
	return nil
}

func StopXray() error {
	return nil
}

// 初始化Xray
func init() {
	xaryPath := constant.XrayPath
	if !util.Exists(xaryPath) {
		if err := os.MkdirAll(xaryPath, os.ModePerm); err != nil {
			logrus.Errorf("创建xray文件夹异常 err: %v\n", err)
			panic(err)
		}
	}

	// 创建默认Xray配置模板文件
	xrayConfigFilePath := constant.XrayFilePath
	if !util.Exists(xrayConfigFilePath) {
		file, err := os.Create(xrayConfigFilePath)
		if err != nil {
			logrus.Errorf("创建xray config.json文件异常 err: %v\n", err)
			panic(err)
		}
		defer file.Close()

		_, err = file.WriteString(`{
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
      "port": 10087,
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
`)
		if err != nil {
			logrus.Errorf("xray config.json文件写入异常 err: %v\n", err)
			panic(err)
		}
	}
}
