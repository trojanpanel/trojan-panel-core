package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xtls/xray-core/common/net"
	"google.golang.org/grpc"
	"trojan-panel-core/api"
	"trojan-panel-core/core"
	"trojan-panel-core/dao"
	"trojan-panel-core/middleware"
	"trojan-panel-core/router"
	"trojan-panel-core/util"
)

func main() {
	rpcServer := grpc.NewServer()
	api.RegisterApiNodeServiceServer(rpcServer, new(api.ServerApi))
	listener, err := net.Listen("tcp", ":8100")
	if err != nil {
		panic(fmt.Sprintf("gRPC服务监听端口失败%v", err))
	}
	_ = rpcServer.Serve(listener)
	r := gin.Default()
	router.Router(r)
	_ = r.Run(":8082")
}

func init() {
	util.InitFile()
	core.InitConfig()
	middleware.InitLog()
	dao.InitMySQL()
	middleware.InitCron()
	middleware.InitRateLimiter()
	api.InitValidator()
}
