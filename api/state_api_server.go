package api

import (
	"context"
)

type StateApiServer struct {
}

func (s *StateApiServer) mustEmbedUnimplementedApiStateServiceServer() {
}

func (s *StateApiServer) Ping(ctx context.Context, stateDto *StateDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: "pong"}, nil
}
