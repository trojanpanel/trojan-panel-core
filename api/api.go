package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/common/net"
	"google.golang.org/grpc"
	"os"
	"trojan-core/api/proxy"
	"trojan-core/api/server"
	"trojan-core/api/version"
	"trojan-core/model/constant"
)

func StarGrpcServer() error {
	rpcServer := grpc.NewServer()
	proxy.RegisterApiProxyServiceServer(rpcServer, new(proxy.ApiProxyService))
	server.RegisterApiServerServiceServer(rpcServer, new(server.ApiServerService))
	version.RegisterApiVersionServiceServer(rpcServer, new(version.ApiVersionService))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv(constant.GrpcPort)))
	if err != nil {
		logrus.Errorf(fmt.Sprintf("gRPC server listening port err: %v", err))
		return fmt.Errorf("gRPC server listening port err")
	}
	return rpcServer.Serve(listener)
}
