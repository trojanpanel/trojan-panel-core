package api

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
	"trojan-core/model/constant"
)

func newGrpcInstance(token, target string, timeout time.Duration) (conn *grpc.ClientConn, ctx context.Context, clo func(), err error) {
	tokenParam := TokenValidateParam{
		Token: token,
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(&tokenParam),
	}
	conn, err = grpc.Dial(target, opts...)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	clo = func() {
		cancel()
		if conn != nil {
			conn.Close()
		}
	}
	if err != nil {
		err = errors.New(constant.GrpcError)
	}
	return
}
