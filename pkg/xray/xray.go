package xray

import (
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/common/cmdarg"
	"os"
	"strings"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/pkg/xray/start"
	"trojan-panel-core/util"
)

// StartXray 启动Xray
func StartXray() {
	os.Args = []string{"run"}
	start.SetConfigFiles(cmdarg.Arg{constant.XrayConfigFilePath})
	start.XrayMain()
}

// StopXray 关闭Xray
func StopXray() error {
	return nil
}

func XrayConfig() {
	// 创建默认Xray配置模板文件
	xrayConfigFilePath := constant.XrayConfigFilePath
	if !util.Exists(xrayConfigFilePath) {
		file, err := os.Create(xrayConfigFilePath)
		if err != nil {
			logrus.Errorf("创建xray config.json文件异常 err: %v\n", err)
			panic(err)
		}
		defer file.Close()

		xrayConfigContent := `{
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
      "port": ${inbounds_api_port},
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
		xrayConfigContent = strings.ReplaceAll(xrayConfigContent, "${inbounds_api_port}", constant.GrpcPortXray)
		_, err = file.WriteString(xrayConfigContent)
		if err != nil {
			logrus.Errorf("xray config.json文件写入异常 err: %v\n", err)
			panic(err)
		}
	}
}
