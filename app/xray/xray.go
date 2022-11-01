package xray

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
	"trojan-panel-core/core"
	"trojan-panel-core/core/process"
	"trojan-panel-core/module/bo"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/util"
)

// InitXrayApp 初始化Xray应用
func InitXrayApp() error {
	apiPorts, err := util.GetConfigApiPorts(constant.XrayPath)
	if err != nil {
		return err
	}
	xrayProcess := process.NewXrayProcess()
	for _, apiPort := range apiPorts {
		// 启动xray
		if err = xrayProcess.StartXray(apiPort); err != nil {
			return err
		}
	}
	return nil
}

// StartXray 启动Xray
func StartXray(xrayConfigDto dto.XrayConfigDto) error {
	var err error
	if err = initXray(xrayConfigDto); err != nil {
		return err
	}
	if err = process.NewXrayProcess().StartXray(xrayConfigDto.ApiPort); err != nil {
		return err
	}
	return nil
}

// StopXray 暂停Xray
func StopXray(apiPort uint, removeFile bool) error {
	if err := process.NewXrayProcess().Stop(apiPort, removeFile); err != nil {
		logrus.Errorf("xray stop err: %v", err)
		return err
	}
	return nil
}

// RestartXray 重启Xray
func RestartXray(apiPort uint) error {
	if err := StopXray(apiPort, false); err != nil {
		return err
	}
	if err := StartXray(dto.XrayConfigDto{
		ApiPort: apiPort,
	}); err != nil {
		return err
	}
	return nil
}

// 初始化Xray文件
func initXray(xrayConfigDto dto.XrayConfigDto) error {
	// 初始化文件夹
	xrayPath := constant.XrayPath
	if !util.Exists(xrayPath) {
		if err := os.MkdirAll(xrayPath, os.ModePerm); err != nil {
			logrus.Errorf("创建Xray文件夹异常 err: %v", err)
			return err
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
			logrus.Errorf("Xray二进制文件下载失败 err: %v", err)
			return err
		}
	}

	// 初始化配置 文件名称格式：config-[apiPort]-[protocol].json
	xrayConfigFilePath := fmt.Sprintf("%s/config-%d-%s.json", constant.XrayPath, xrayConfigDto.ApiPort, xrayConfigDto.Protocol)
	file, err := os.OpenFile(xrayConfigFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		logrus.Errorf("创建xray %s文件异常 err: %v", xrayConfigFilePath, err)
		return err
	}
	defer file.Close()

	// 根据不同的协议生成对应的配置文件，用户信息通过新建同步协程
	configTemplateContent := `{
    "log": {
        "loglevel": "warning"
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
            "protocol": "freedom"
        }
    ],
    "api": {
        "tag": "api",
        "services": [
            "HandlerService",
            "LoggerService",
            "StatsService"
        ]
    },
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
    },
    "stats": {},
    "policy": {
        "levels": {
            "0": {
                "statsUserUplink": true,
                "statsUserDownlink": true
            }
        },
        "system": {
            "statsInboundUplink": true,
            "statsInboundDownlink": true
        }
    }
}
`
	configTemplateContent = strings.ReplaceAll(configTemplateContent, "${api_port}", strconv.FormatInt(int64(xrayConfigDto.ApiPort), 10))
	xrayConfig := &bo.XrayConfigBo{}
	// 将json字符串映射到模板对象
	if err = json.Unmarshal([]byte(configTemplateContent), xrayConfig); err != nil {
		logrus.Errorf("xray template config反序列化异常 err: %v", err)
		return err
	}

	// 设置streamSettings字段
	streamSettingsStr := []byte("{}")
	if xrayConfigDto.StreamSettings != "" {
		streamSettings := &bo.StreamSettings{}
		if err = json.Unmarshal([]byte(xrayConfigDto.StreamSettings), streamSettings); err != nil {
			logrus.Errorf("xray StreamSettings反序列化异常 err: %v", err)
			return err
		}

		if streamSettings.Security != "none" {
			// 设置证书
			certConfig := core.Config.CertConfig
			var certificates []bo.Certificate
			certificate := bo.Certificate{
				CertificateFile: certConfig.CrtPath,
				KeyFile:         certConfig.KeyPath,
			}
			certificates = append(certificates, certificate)
			if streamSettings.Security == "tls" {
				streamSettings.TlsSettings.Certificates = certificates
			} else if streamSettings.Security == "xtls" {
				streamSettings.XtlsSettings.Certificates = certificates
			}
		}

		streamSettingsStr, err = json.MarshalIndent(streamSettings, "", "    ")
		if err != nil {
			logrus.Errorf("xray StreamSettings序列化异常 err: %v", err)
			return err
		}
	}

	// 添加入站协议
	xrayConfig.Inbounds = append(xrayConfig.Inbounds, bo.InboundBo{
		Listen:         "0.0.0.0",
		Port:           xrayConfigDto.Port,
		Protocol:       xrayConfigDto.Protocol,
		Settings:       bo.TypeMessage(xrayConfigDto.Settings),
		StreamSettings: streamSettingsStr,
		Tag:            xrayConfigDto.Tag,
		Sniffing:       bo.TypeMessage(xrayConfigDto.Sniffing),
		Allocate:       bo.TypeMessage(xrayConfigDto.Allocate),
	})
	configContentByte, err := json.MarshalIndent(xrayConfig, "", "    ")
	if err != nil {
		logrus.Errorf("xray template config反序列化异常 err: %v", err)
		return err
	}
	_, err = file.Write(configContentByte)
	if err != nil {
		logrus.Errorf("xray config.json文件写入异常 err: %v", err)
		return err
	}
	return nil
}
