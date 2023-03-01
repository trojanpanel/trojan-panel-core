package api

import (
	"context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"trojan-panel-core/core/process"
	"trojan-panel-core/module/constant"
)

type StateApiServer struct {
}

func (s *StateApiServer) mustEmbedUnimplementedApiStateServiceServer() {
}

func (s *StateApiServer) Ping(ctx context.Context, stateDto *StateDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	state := process.GetState(uint(stateDto.GetNodeTypeId()), uint(stateDto.GetPort())+30000)
	stateVo := &StateVo{
		NodeState: 0,
		Version:   constant.TrojanPanelCoreVersion,
	}
	if state {
		stateVo.NodeState = 1
	}
	data, err := anypb.New(proto.Message(stateVo))
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: "pong", Data: data}, nil
}
