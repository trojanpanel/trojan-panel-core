package api

import (
	"context"
	"errors"
	"trojan-panel-core/app"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/util"
)

type NodeServerApi struct {
}

func (s *NodeServerApi) mustEmbedUnimplementedApiNodeServiceServer() {
}

func (s *NodeServerApi) AddNode(ctx context.Context, nodeAddDto *NodeAddDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	var err error
	if (nodeAddDto.XrayPort != 0 && (nodeAddDto.XrayPort <= 100 || nodeAddDto.XrayPort >= 30000)) ||
		(nodeAddDto.TrojanGoPort != 0 && (nodeAddDto.TrojanGoPort <= 100 || nodeAddDto.TrojanGoPort >= 30000)) ||
		(nodeAddDto.HysteriaPort != 0 && (nodeAddDto.HysteriaPort <= 100 || nodeAddDto.HysteriaPort >= 30000)) {
		err = errors.New(constant.PortRangeError)
	} else {
		if nodeAddDto.XrayPort != 0 {
			if !util.IsPortAvailable(uint(nodeAddDto.XrayPort), "tcp") {
				err = errors.New(constant.PortIsOccupied)
			}
			if !util.IsPortAvailable(uint(nodeAddDto.XrayPort+10000), "tcp") {
				err = errors.New(constant.PortIsOccupied)
			}
		} else if nodeAddDto.TrojanGoPort != 0 {
			if !util.IsPortAvailable(uint(nodeAddDto.TrojanGoPort), "tcp") {
				err = errors.New(constant.PortIsOccupied)
			}
			if !util.IsPortAvailable(uint(nodeAddDto.TrojanGoPort+10000), "tcp") {
				err = errors.New(constant.PortIsOccupied)
			}
		} else if nodeAddDto.HysteriaPort != 0 {
			if !util.IsPortAvailable(uint(nodeAddDto.HysteriaPort), "udp") {
				err = errors.New(constant.PortIsOccupied)
			}
		}
	}
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	if err := app.StartApp(dto.NodeAddDto{
		NodeTypeId: uint(nodeAddDto.NodeTypeId),
		// Xray
		XrayPort:           uint(nodeAddDto.XrayPort),
		XrayProtocol:       nodeAddDto.XrayProtocol,
		XraySettings:       nodeAddDto.XraySettings,
		XrayStreamSettings: nodeAddDto.XrayStreamSettings,
		XrayTag:            nodeAddDto.XrayTag,
		XraySniffing:       nodeAddDto.XraySniffing,
		XrayAllocate:       nodeAddDto.XrayAllocate,
		// Trojan Go
		TrojanGoPort:            uint(nodeAddDto.TrojanGoPort),
		TrojanGoIp:              nodeAddDto.TrojanGoIp,
		TrojanGoSni:             nodeAddDto.TrojanGoSni,
		TrojanGoMuxEnable:       uint(nodeAddDto.TrojanGoMuxEnable),
		TrojanGoWebsocketEnable: uint(nodeAddDto.TrojanGoWebsocketEnable),
		TrojanGoWebsocketPath:   nodeAddDto.TrojanGoWebsocketPath,
		TrojanGoWebsocketHost:   nodeAddDto.TrojanGoWebsocketHost,
		TrojanGoSSEnable:        uint(nodeAddDto.TrojanGoSSEnable),
		TrojanGoSSMethod:        nodeAddDto.TrojanGoSSMethod,
		TrojanGoSSPassword:      nodeAddDto.TrojanGoSSPassword,
		// Hysteria
		HysteriaPort:     uint(nodeAddDto.HysteriaPort),
		HysteriaProtocol: nodeAddDto.HysteriaProtocol,
		HysteriaIp:       nodeAddDto.HysteriaIp,
		HysteriaUpMbps:   int(nodeAddDto.HysteriaUpMbps),
		HysteriaDownMbps: int(nodeAddDto.HysteriaDownMbps),
	}); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}

func (s *NodeServerApi) RemoveNode(ctx context.Context, nodeRemoveDto *NodeRemoveDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	if err := app.StopApp(uint(nodeRemoveDto.Port)+100, uint(nodeRemoveDto.NodeType)); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}
