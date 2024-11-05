package api

import (
	"context"
)

type NodeApi struct {
}

func (s *NodeApi) CreateNode(ctx context.Context, nodeCreateDto *NodeCreateDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}

func (s *NodeApi) DeleteNode(ctx context.Context, nodeDeleteDto *NodeDeleteDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}

func (s *NodeApi) GetNode(ctx context.Context, nodeDto *NodeDto) (*Response, error) {
	if err := authRequest(ctx); err != nil {
		return &Response{Success: false, Msg: err.Error()}, nil
	}
	return &Response{Success: true, Msg: ""}, nil
}
