package api

import (
	"context"
	"trojan-panel-core/app"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
)

type NodeServerApi struct {
}

func (s *NodeServerApi) mustEmbedUnimplementedApiNodeServiceServer() {
}

func (s *NodeServerApi) AddNode(ctx context.Context, nodeAddDto *NodeAddDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}

	// 校验端口
	if nodeAddDto.Port != 0 && (nodeAddDto.Port <= 100 || nodeAddDto.Port >= 30000) {
		return &Response{Success: false, Msg: constant.PortRangeError}, nil
	}

	if err := app.StartApp(dto.NodeAddDto{
		NodeTypeId: uint(nodeAddDto.NodeTypeId),
		Port:       uint(nodeAddDto.Port),

		// Xray
		XrayProtocol:       nodeAddDto.XrayProtocol,
		XraySettings:       nodeAddDto.XraySettings,
		XrayStreamSettings: nodeAddDto.XrayStreamSettings,
		XrayTag:            nodeAddDto.XrayTag,
		XraySniffing:       nodeAddDto.XraySniffing,
		XrayAllocate:       nodeAddDto.XrayAllocate,
		// Trojan Go
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
	if err := app.StopApp(uint(nodeRemoveDto.Port)+10000, uint(nodeRemoveDto.NodeType)); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}
