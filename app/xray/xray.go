package xray

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"trojan-panel-core/core"
	"trojan-panel-core/core/process"
	"trojan-panel-core/model/bo"
	"trojan-panel-core/model/constant"
	"trojan-panel-core/model/dto"
	"trojan-panel-core/util"
)

func InitXrayApp() error {
	apiPorts, err := util.GetConfigApiPorts(constant.XrayPath)
	if err != nil {
		return err
	}
	xrayProcess := process.NewXrayProcess()
	for _, apiPort := range apiPorts {
		if err = xrayProcess.StartXray(apiPort); err != nil {
			return err
		}
	}
	return nil
}

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

func StopXray(apiPort uint, removeFile bool) error {
	if err := process.NewXrayProcess().Stop(apiPort, removeFile); err != nil {
		logrus.Errorf("xray stop err: %v", err)
		return err
	}
	return nil
}

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

func initXray(xrayConfigDto dto.XrayConfigDto) error {
	// initialization configuration file name format: config-[apiPort]-[protocol].json
	xrayConfigFilePath := fmt.Sprintf("%s/config-%d-%s.json", constant.XrayPath, xrayConfigDto.ApiPort, xrayConfigDto.Protocol)
	file, err := os.OpenFile(xrayConfigFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		logrus.Errorf("create xray file %s err: %v", xrayConfigFilePath, err)
		return err
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	// generate corresponding configuration files according to different protocols, and account information is created through a new synchronous coroutine
	if xrayConfigDto.Template == "" {
		xrayConfigDto.Template = `{
    "log": {
        "loglevel": "warning"
    },
    "inbounds": [],
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
}`
	}
	xrayConfig := &bo.XrayConfigBo{}
	// map json strings to template objects
	if err = json.Unmarshal([]byte(xrayConfigDto.Template), xrayConfig); err != nil {
		logrus.Errorf("xray template config deserialization err: %v", err)
		return err
	}

	// set the streamSettings field
	streamSettingsStr := []byte("{}")
	if xrayConfigDto.StreamSettings != "" {
		streamSettings := &bo.StreamSettings{}
		if err = json.Unmarshal([]byte(xrayConfigDto.StreamSettings), streamSettings); err != nil {
			logrus.Errorf("xray StreamSettings deserialization err: %v", err)
			return err
		}

		if streamSettings.Security != "none" {
			// set cert
			certConfig := core.Config.CertConfig
			var certificates []bo.Certificate
			certificate := bo.Certificate{
				CertificateFile: certConfig.CrtPath,
				KeyFile:         certConfig.KeyPath,
			}
			certificates = append(certificates, certificate)
			if streamSettings.Security == "tls" && len(streamSettings.TlsSettings.Certificates) == 0 {
				streamSettings.TlsSettings.Certificates = certificates
			} else if streamSettings.Security == "reality" {

			}
		}

		streamSettingsStr, err = json.MarshalIndent(streamSettings, "", "    ")
		if err != nil {
			logrus.Errorf("xray StreamSettings serialization err: %v", err)
			return err
		}
	}

	// add inbound protocol
	xrayConfig.Inbounds = append(xrayConfig.Inbounds, bo.InboundBo{
		Listen:   "127.0.0.1",
		Port:     xrayConfigDto.ApiPort,
		Protocol: "dokodemo-door",
		Settings: bo.TypeMessage("{\"address\": \"127.0.0.1\"}"),
		Tag:      "api",
	})

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
		logrus.Errorf("xray template config deserialization err: %v", err)
		return err
	}
	_, err = file.Write(configContentByte)
	if err != nil {
		logrus.Errorf("xray file config.json write err: %v", err)
		return err
	}
	return nil
}

func InitXrayBinFile() error {
	xrayPath := constant.XrayPath
	if !util.Exists(xrayPath) {
		if err := os.MkdirAll(xrayPath, os.ModePerm); err != nil {
			logrus.Errorf("create xray folder err: %v", err)
			return err
		}
	}

	binaryFilePath, err := util.GetBinaryFilePath(constant.Xray)
	if err != nil {
		return err
	}
	if !util.Exists(binaryFilePath) {
		logrus.Errorf("xray binary does not exist")
		return errors.New(constant.BinaryFileNotExist)
	}
	return nil
}
