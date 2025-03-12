package version

import (
	"context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"trojan-core/model/constant"
	"trojan-core/util"
)

type ApiVersionService struct {
}

func (a *ApiVersionService) GetVersion(ctx context.Context, apiVersionDto *ApiVersionDto) (*Response, error) {
	if err := util.AuthRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	apiVersionVo := &ApiVersionVo{
		SystemVersion: constant.SystemVersion,
	}
	data, err := anypb.New(proto.Message(apiVersionVo))
	if err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: "", Data: data}, nil
}
