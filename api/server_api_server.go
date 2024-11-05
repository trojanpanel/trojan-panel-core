package api

import (
	"context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"trojan-core/model/constant"
	"trojan-core/util"
)

type ServerApi struct {
}

func (s *ServerApi) GetServer(ctx context.Context, serverDto *ServerDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	cpuUsed, err := util.GetCpuPercent()
	memUsed, err := util.GetMemPercent()
	diskUsed, err := util.GetDiskPercent()
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	serverVo := &ServerVo{
		CpuUsed:  float32(cpuUsed),
		MemUsed:  float32(memUsed),
		DiskUsed: float32(diskUsed),
		Version:  constant.Version,
	}
	data, err := anypb.New(proto.Message(serverVo))
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: "", Data: data}, nil
}
