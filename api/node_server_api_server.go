package api

import (
	"context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"trojan-panel-core/util"
)

type NodeServerApiServer struct {
}

func (s *NodeServerApiServer) mustEmbedUnimplementedApiNodeServerServiceServer() {
}

func (s *NodeServerApiServer) NodeServerState(ctx context.Context, nodeServerGroupDto *NodeServerGroupDto) (*Response, error) {
	//if err := authRequest(ctx); err != nil {
	//	return &Response{Success: false, Msg: err.Error()}, nil
	//}
	cpuUsed, err := util.GetCpuPercent()
	memUsed, err := util.GetMemPercent()
	diskUsed, err := util.GetDiskPercent()
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	nodeServerGroupVo := &NodeServerGroupVo{
		CpuUsed:  float32(cpuUsed),
		MemUsed:  float32(memUsed),
		DiskUsed: float32(diskUsed),
	}
	data, err := anypb.New(proto.Message(nodeServerGroupVo))
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: "", Data: data}, nil
}
