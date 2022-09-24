package app

import (
	"errors"
	"trojan-panel-core/app/hysteria"
	"trojan-panel-core/app/trojango"
	"trojan-panel-core/app/xray"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/util"
)

func StartApp(nodeAddDto dto.NodeAddDto) error {
	port, err := util.GetPortAvailBetween()
	if err != nil {
		return err
	}
	switch nodeAddDto.NodeTypeId {
	case 1:
		if err = xray.StartXray(dto.XrayConfigDto{
			ApiPort:        port + 100,
			Port:           port,
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
		if err = trojango.StartTrojanGo(dto.TrojanGoConfigDto{
			ApiPort:         port + 100,
			Port:            port,
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
		if err = hysteria.StartHysteria(dto.HysteriaConfigDto{
			ApiPort:  port + 100,
			Port:     port,
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
