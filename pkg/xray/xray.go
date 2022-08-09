package xray

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"trojan-panel-core/core/process"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/util"
)

var xrayProcess *process.XrayProcess

// StartXray 启动Xray
func StartXray(xrayConfigDto dto.XrayConfigDto) error {
	var err error
	if err = initXray(xrayConfigDto); err != nil {
		return err
	}
	xrayProcess, err = process.NewXrayProcess(xrayConfigDto.ApiPort)
	if err != nil {
		return err
	}
	if err = xrayProcess.StartXray(xrayConfigDto.ApiPort); err != nil {
		return err
	}
	return nil
}

// StopXray 关闭Xray
func StopXray(apiPort string) error {
	if xrayProcess != nil {
		if err := xrayProcess.Stop(apiPort); err != nil {
			return err
		}
	}
	return nil
}

func initXray(xrayConfigDto dto.XrayConfigDto) error {
	// 初始化文件夹
	xrayPath := constant.XrayPath
	if !util.Exists(xrayPath) {
		if err := os.MkdirAll(xrayPath, os.ModePerm); err != nil {
			logrus.Errorf("创建Xray文件夹异常 err: %v\n", err)
			panic(err)
		}
	}

	// 下载二进制文件
	binaryFilePath, err := util.GetBinaryFilePath(1)
	if err != nil {
		return err
	}
	if !util.Exists(binaryFilePath) {
		if err = util.DownloadFile(fmt.Sprintf("%s/xray-%s-%s", constant.DownloadBaseUrl, runtime.GOOS, runtime.GOARCH),
			binaryFilePath); err != nil {
			logrus.Errorf("Xray二进制文件下载失败 err: %v\n", err)
			panic(err)
		}
	}

	// 初始化配置
	xrayConfigFilePath, err := util.GetConfigFilePath(1, xrayConfigDto.ApiPort)
	if err != nil {
		return err
	}
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
		configContent = strings.ReplaceAll(configContent, "${api_port}", xrayConfigDto.ApiPort)
		_, err = file.WriteString(configContent)
		if err != nil {
			logrus.Errorf("xray config.json文件写入异常 err: %v\n", err)
			return err
		}
	}
	return nil
}
