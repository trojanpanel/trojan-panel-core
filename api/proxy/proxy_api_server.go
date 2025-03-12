package proxy

import (
	"context"
	"trojan-core/util"
)

type ApiProxyService struct {
}

func (a *ApiProxyService) StartProxy(ctx context.Context, startProxyDto *StartProxyDto) (*Response, error) {
	if err := util.AuthRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}

func (a *ApiProxyService) StopProxy(ctx context.Context, stopProxyDto *StopProxyDto) (*Response, error) {
	if err := util.AuthRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}
