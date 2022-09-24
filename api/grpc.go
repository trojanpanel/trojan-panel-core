package api

import (
	"fmt"
	"github.com/xtls/xray-core/common/net"
	"google.golang.org/grpc"
)

func InitGrpcServer() {
	go func() {
		rpcServer := grpc.NewServer()
		RegisterApiNodeServiceServer(rpcServer, new(ServerApi))
		listener, err := net.Listen("tcp", ":8100")
		if err != nil {
			panic(fmt.Sprintf("gRPC服务监听端口失败%v", err))
		}
		_ = rpcServer.Serve(listener)
	}()
}
