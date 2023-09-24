package api

import (
	"context"
	"errors"
	"trojan-panel-core/app"
	"trojan-panel-core/model/constant"
	"trojan-panel-core/model/dto"
	"trojan-panel-core/util"
)

type NodeApiServer struct {
}

func (s *NodeApiServer) AddNode(ctx context.Context, nodeAddDto *NodeAddDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}

	// check port
	var err error
	if nodeAddDto.Port != 0 && (nodeAddDto.Port <= 100 || nodeAddDto.Port >= 30000) {
		err = errors.New(constant.PortRangeError)
	}
	if nodeAddDto.NodeTypeId == constant.Xray || nodeAddDto.NodeTypeId == constant.TrojanGo || nodeAddDto.NodeTypeId == constant.NaiveProxy {
		if !util.IsPortAvailable(uint(nodeAddDto.Port), "tcp") {
			err = errors.New(constant.PortIsOccupied)
		}
		if !util.IsPortAvailable(uint(nodeAddDto.Port+30000), "tcp") {
			err = errors.New(constant.PortIsOccupied)
		}
	} else if nodeAddDto.NodeTypeId == constant.Hysteria || nodeAddDto.NodeTypeId == constant.Hysteria2 {
		if !util.IsPortAvailable(uint(nodeAddDto.Port), "udp") {
			err = errors.New(constant.PortIsOccupied)
		}
	}
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}

	if err = app.StartApp(dto.NodeAddDto{
		NodeTypeId: uint(nodeAddDto.NodeTypeId),
		Port:       uint(nodeAddDto.Port),
		Domain:     nodeAddDto.Domain,

		// Xray
		XrayTemplate:       nodeAddDto.XrayTemplate,
		XrayFlow:           nodeAddDto.XrayFlow,
		XraySSMethod:       nodeAddDto.XraySSMethod,
		XrayProtocol:       nodeAddDto.XrayProtocol,
		XraySettings:       nodeAddDto.XraySettings,
		XrayStreamSettings: nodeAddDto.XrayStreamSettings,
		XrayTag:            nodeAddDto.XrayTag,
		XraySniffing:       nodeAddDto.XraySniffing,
		XrayAllocate:       nodeAddDto.XrayAllocate,
		// Trojan Go
		TrojanGoSni:             nodeAddDto.TrojanGoSni,
		TrojanGoMuxEnable:       uint(nodeAddDto.TrojanGoMuxEnable),
		TrojanGoWebsocketEnable: uint(nodeAddDto.TrojanGoWebsocketEnable),
		TrojanGoWebsocketPath:   nodeAddDto.TrojanGoWebsocketPath,
		TrojanGoWebsocketHost:   nodeAddDto.TrojanGoWebsocketHost,
		TrojanGoSSEnable:        uint(nodeAddDto.TrojanGoSSEnable),
		TrojanGoSSMethod:        nodeAddDto.TrojanGoSSMethod,
		TrojanGoSSPassword:      nodeAddDto.TrojanGoSSPassword,
		// Hysteria
		HysteriaProtocol: nodeAddDto.HysteriaProtocol,
		HysteriaObfs:     nodeAddDto.HysteriaObfs,
		HysteriaUpMbps:   int(nodeAddDto.HysteriaUpMbps),
		HysteriaDownMbps: int(nodeAddDto.HysteriaDownMbps),
		// Hysteria2
		Hysteria2ObfsPassword: nodeAddDto.Hysteria2ObfsPassword,
		Hysteria2UpMbps:       int(nodeAddDto.Hysteria2UpMbps),
		Hysteria2DownMbps:     int(nodeAddDto.Hysteria2DownMbps),
	}); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}

func (s *NodeApiServer) RemoveNode(ctx context.Context, nodeRemoveDto *NodeRemoveDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	if err := app.StopApp(uint(nodeRemoveDto.Port)+30000, uint(nodeRemoveDto.NodeTypeId)); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}
