package xray

import (
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/common/cmdarg"
	"os"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/pkg/xray/start"
	"trojan-panel-core/util"
)

// StartXray 启动Xray
func StartXray() {
	start.SetConfigFiles(cmdarg.Arg{constant.XrayConfigFilePath})
	start.XrayMain()
}

// StopXray 关闭Xray
func StopXray() error {
	return nil
}

// 初始化Xray
func init() {
	// 创建默认Xray配置模板文件
	xrayConfigFilePath := constant.XrayConfigFilePath
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
