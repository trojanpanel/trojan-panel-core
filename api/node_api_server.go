package api

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"trojan-panel-core/app"
	"trojan-panel-core/dao/redis"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/util"
)

type ServerApi struct {
}

func (s *ServerApi) mustEmbedUnimplementedApiNodeServiceServer() {
}

func (s *ServerApi) AddNode(ctx context.Context, nodeAddDto *NodeAddDto) (*NodeResponse, error) {
	if err := authRequest(ctx); err != nil {
		return &NodeResponse{Success: false, Msg: err.Error()}, nil
	}
	if err := app.StartApp(dto.NodeAddDto{
		NodeType:                uint(nodeAddDto.NodeType),
		XrayPort:                uint(nodeAddDto.XrayPort),
		XrayProtocol:            nodeAddDto.XrayProtocol,
		XraySettings:            nodeAddDto.XraySettings,
		XrayStreamSettings:      nodeAddDto.XrayStreamSettings,
		XrayTag:                 nodeAddDto.XrayTag,
		XraySniffing:            nodeAddDto.XraySniffing,
		XrayAllocate:            nodeAddDto.XrayAllocate,
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
		HysteriaPort:            uint(nodeAddDto.HysteriaPort),
		HysteriaProtocol:        nodeAddDto.HysteriaProtocol,
		HysteriaIp:              nodeAddDto.HysteriaIp,
		HysteriaUpMbps:          int(nodeAddDto.HysteriaUpMbps),
		HysteriaDownMbps:        int(nodeAddDto.HysteriaDownMbps),
	}); err != nil {
		return &NodeResponse{Success: false, Msg: err.Error()}, nil
	}
	return &NodeResponse{Success: true, Msg: ""}, nil
}

func (s *ServerApi) RemoveNode(ctx context.Context, nodeRemoveDto *NodeRemoveDto) (*NodeResponse, error) {
	if err := authRequest(ctx); err != nil {
		return &NodeResponse{Success: false, Msg: err.Error()}, nil
	}
	if err := app.StopApp(uint(nodeRemoveDto.Port)+100, uint(nodeRemoveDto.NodeType)); err != nil {
		return &NodeResponse{Success: false, Msg: err.Error()}, nil
	}
	return &NodeResponse{Success: true, Msg: ""}, nil
}

// token认证
func authRequest(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New(constant.UnauthorizedError)
	}
	var token string
	if val, ok := md["token"]; ok {
		token = val[0]
	}
	myClaims, err := util.ParseToken(token)
	if err != nil {
		return errors.New(constant.UnauthorizedError)
	}
	get := redis.Client.String.
		Get(fmt.Sprintf("trojan-panel:token:%s", myClaims.AccountVo.Username))
	result, err := get.String()
	if err != nil || result == "" {
		return errors.New(constant.IllegalTokenError)
	}
	return nil
}
