package api

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/common/net"
	"google.golang.org/grpc"
	"os"
	"trojan-core/api/proxy"
	"trojan-core/api/server"
	"trojan-core/api/version"
	"trojan-core/middleware"
	"trojan-core/model/constant"
	"trojan-core/util"
)

func StarServer() error {
	defer releaseResource()

	if err := initFile(); err != nil {
		return err
	}
	if err := middleware.InitCron(); err != nil {
		return err
	}

	rpcServer := grpc.NewServer()
	proxy.RegisterApiProxyServiceServer(rpcServer, new(proxy.ApiProxyService))
	server.RegisterApiServerServiceServer(rpcServer, new(server.ApiServerService))
	version.RegisterApiVersionServiceServer(rpcServer, new(version.ApiVersionService))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv(constant.GrpcPort)))
	if err != nil {
		logrus.Errorf(fmt.Sprintf("gRPC server listening port err: %v", err))
		return errors.New("gRPC server listening port err")
	}
	return rpcServer.Serve(listener)
}

func releaseResource() {

}

func initFile() error {
	var dirs = []string{constant.LogDir, constant.BinDir,
		constant.XrayConfigDir, constant.SingBoxConfigDir,
		constant.HysteriaConfigDir, constant.NaiveProxyConfigDir,
	}
	for _, item := range dirs {
		if !util.Exists(item) {
			if err := os.Mkdir(item, os.ModePerm); err != nil {
				logrus.Errorf("%s create err: %v", item, err)
				return errors.New(fmt.Sprintf("%s create err", item))
			}
		}
	}
	return nil
}
