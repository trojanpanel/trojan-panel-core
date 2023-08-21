package api

import (
	"fmt"
	"github.com/xtls/xray-core/common/net"
	"google.golang.org/grpc"
	"trojan-panel-core/core"
)

func InitGrpcServer() {
	go func() {
		grpcConfig := core.Config.GrpcConfig
		rpcServer := grpc.NewServer()
		RegisterApiNodeServiceServer(rpcServer, new(NodeApiServer))
		RegisterApiAccountServiceServer(rpcServer, new(AccountApiServer))
		RegisterApiStateServiceServer(rpcServer, new(StateApiServer))
		RegisterApiNodeServerServiceServer(rpcServer, new(NodeServerApiServer))
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcConfig.Port))
		if err != nil {
			panic(fmt.Sprintf("gRPC service listening port err: %v", err))
		}
		_ = rpcServer.Serve(listener)
	}()
}
