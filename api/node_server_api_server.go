package api

import (
	"context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"trojan-panel-core/util"
)

type NodeServerApiServer struct {
}

func (s *NodeServerApiServer) GetNodeServerInfo(ctx context.Context, nodeServerInfoDto *NodeServerInfoDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	cpuUsed, err := util.GetCpuPercent()
	memUsed, err := util.GetMemPercent()
	diskUsed, err := util.GetDiskPercent()
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	nodeServerInfoVo := &NodeServerInfoVo{
		CpuUsed:  float32(cpuUsed),
		MemUsed:  float32(memUsed),
		DiskUsed: float32(diskUsed),
	}
	data, err := anypb.New(proto.Message(nodeServerInfoVo))
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: "", Data: data}, nil
}
