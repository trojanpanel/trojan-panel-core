package api

import (
	"context"
)

type TokenValidateParam struct {
	Token string `json:"token"`
}

func (x *TokenValidateParam) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"token": x.Token,
	}, nil
}

func (x *TokenValidateParam) RequireTransportSecurity() bool {
	return false
}
