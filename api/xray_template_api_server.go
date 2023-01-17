package api

import (
	"context"
	"trojan-panel-core/app/xray"
)

type XrayTemplateApiServer struct {
}

func (s *XrayTemplateApiServer) mustEmbedUnimplementedApiXrayTemplateServiceServer() {
}

func (s *XrayTemplateApiServer) UpdateXrayTemplate(ctx context.Context, xrayTemplateDto *XrayTemplateDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	if err := xray.UpdateXrayTemplate(xrayTemplateDto.XrayTemplate); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}
