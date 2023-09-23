package api

import (
	"context"
	"trojan-panel-core/app"
)

type AccountApiServer struct {
}

func (s *AccountApiServer) RemoveAccount(ctx context.Context, accountRemoveDto *AccountRemoveDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	if err := app.RemoveAccount(accountRemoveDto.Password); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}
