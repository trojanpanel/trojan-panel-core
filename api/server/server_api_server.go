package server

import (
	"context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"trojan-core/service"
	"trojan-core/util"
)

type ApiServerService struct {
}

func (a *ApiServerService) GetServerStats(ctx context.Context, apiServerDto *ApiServerDto) (*Response, error) {
	if err := util.AuthRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	cpuUsed, memUsed, diskUsed, err := service.GetServerStats()
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	apiServerVo := &ApiServerVo{
		CpuUsed:  float32(cpuUsed),
		MemUsed:  float32(memUsed),
		DiskUsed: float32(diskUsed),
	}
	data, err := anypb.New(proto.Message(apiServerVo))
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: "", Data: data}, nil
}
