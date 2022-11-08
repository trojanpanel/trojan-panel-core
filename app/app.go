package app

import (
	"errors"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/app/hysteria"
	"trojan-panel-core/app/trojango"
	"trojan-panel-core/app/xray"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
)

func StartApp(nodeAddDto dto.NodeAddDto) error {
	switch nodeAddDto.NodeTypeId {
	case 1:
		if err := xray.StartXray(dto.XrayConfigDto{
			ApiPort:        nodeAddDto.Port + 10000,
			Port:           nodeAddDto.Port,
			Protocol:       nodeAddDto.XrayProtocol,
			Settings:       nodeAddDto.XraySettings,
			StreamSettings: nodeAddDto.XrayStreamSettings,
			Tag:            nodeAddDto.XrayTag,
			Sniffing:       nodeAddDto.XraySniffing,
			Allocate:       nodeAddDto.XrayAllocate,
		}); err != nil {
			return err
		}
	case 2:
		if err := trojango.StartTrojanGo(dto.TrojanGoConfigDto{
			ApiPort:         nodeAddDto.Port + 10000,
			Port:            nodeAddDto.Port,
			Ip:              nodeAddDto.TrojanGoIp,
			Sni:             nodeAddDto.TrojanGoSni,
			MuxEnable:       nodeAddDto.TrojanGoMuxEnable,
			WebsocketEnable: nodeAddDto.TrojanGoWebsocketEnable,
			WebsocketPath:   nodeAddDto.TrojanGoWebsocketPath,
			WebsocketHost:   nodeAddDto.TrojanGoWebsocketHost,
			SSEnable:        nodeAddDto.TrojanGoSSEnable,
			SSMethod:        nodeAddDto.TrojanGoSSMethod,
			SSPassword:      nodeAddDto.TrojanGoSSPassword,
		}); err != nil {
			return err
		}
	case 3:
		if err := hysteria.StartHysteria(dto.HysteriaConfigDto{
			ApiPort:  nodeAddDto.Port + 10000,
			Port:     nodeAddDto.Port,
			Protocol: nodeAddDto.HysteriaProtocol,
			Ip:       nodeAddDto.HysteriaIp,
			UpMbps:   nodeAddDto.HysteriaUpMbps,
			DownMbps: nodeAddDto.HysteriaDownMbps,
		}); err != nil {
			return err
		}
	default:
		return errors.New(constant.NodeTypeNotExist)
	}
	return nil
}

func StopApp(apiPort uint, nodeType uint) error {
	switch nodeType {
	case 1:
		if err := xray.StopXray(apiPort, true); err != nil {
			return err
		}
	case 2:
		if err := trojango.StopTrojanGo(apiPort, true); err != nil {
			return err
		}
	case 3:
		if err := hysteria.StopHysteria(apiPort, true); err != nil {
			return err
		}
	default:
		return errors.New(constant.NodeTypeNotExist)
	}
	return nil
}

func RestartApp(apiPort uint, nodeType uint) error {
	switch nodeType {
	case 1:
		if err := xray.RestartXray(apiPort); err != nil {
			return err
		}
	case 2:
		if err := trojango.RestartTrojanGo(apiPort); err != nil {
			return err
		}
	case 3:
		if err := hysteria.RestartHysteria(apiPort); err != nil {
			return err
		}
	default:
		return errors.New(constant.NodeTypeNotExist)
	}
	return nil
}

func InitApp() {
	InitBinFile()
	if err := xray.InitXrayApp(); err != nil {
		logrus.Errorf("Xray app 初始化失败 err: %s", err.Error())
	}
	if err := trojango.InitTrojanGoApp(); err != nil {
		logrus.Errorf("TrojanGo app 初始化失败 err: %s", err.Error())
	}
	if err := hysteria.InitHysteriaApp(); err != nil {
		logrus.Errorf("Hysteria app 初始化失败 err: %s", err.Error())
	}
}

func InitBinFile() {
	if err := xray.InitXrayBinFile(); err != nil {
		logrus.Errorf("下载Xray文件异常 err: %v", err)
		panic(err)
	}
	if err := trojango.InitTrojanGoBinFile(); err != nil {
		logrus.Errorf("下载TrojanGo文件异常 err: %v", err)
		panic(err)
	}
	if err := hysteria.InitHysteriaBinFile(); err != nil {
		logrus.Errorf("下载Hysteria文件异常 err: %v", err)
		panic(err)
	}
}
