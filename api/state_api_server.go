package api

import (
	"context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"trojan-panel-core/core/process"
	"trojan-panel-core/model/constant"
)

type StateApiServer struct {
}

func (s *StateApiServer) GetNodeState(ctx context.Context, nodeStateDto *NodeStateDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	if nodeStateDto.GetNodeTypeId() <= 0 || nodeStateDto.GetPort() <= 0 {
		return &Response{Success: false, Msg: "parameter validation error", Data: nil}, nil
	}
	nodeStateVo := NodeStateVo{}
	success := process.GetState(uint(nodeStateDto.GetNodeTypeId()), uint(nodeStateDto.GetPort())+30000)
	if success {
		nodeStateVo.Status = 1
	} else {
		nodeStateVo.Status = 0
	}
	data, err := anypb.New(proto.Message(&nodeStateVo))
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: "", Data: data}, nil
}

func (s *StateApiServer) GetNodeServerState(ctx context.Context, nodeServerStateDto *NodeServerStateDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	nodeServerStateVo := &NodeServerStateVo{
		Version: constant.TrojanPanelCoreVersion,
	}
	data, err := anypb.New(proto.Message(nodeServerStateVo))
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: "", Data: data}, nil
}
