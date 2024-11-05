package api

import (
	"fmt"
	"github.com/xtls/xray-core/common/net"
	"google.golang.org/grpc"
)

func InitGrpcServer(port string) {
	go func() {
		rpcServer := grpc.NewServer()
		RegisterApiNodeServiceServer(rpcServer, new(NodeApi))
		RegisterApiServerServiceServer(rpcServer, new(ServerApi))
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			panic(fmt.Sprintf("gRPC service listening port err: %v", err))
		}
		_ = rpcServer.Serve(listener)
	}()
}
