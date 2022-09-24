package main

import (
	"github.com/gin-gonic/gin"
	"trojan-panel-core/api"
	"trojan-panel-core/core"
	"trojan-panel-core/dao"
	"trojan-panel-core/middleware"
	"trojan-panel-core/router"
	"trojan-panel-core/util"
)

func main() {
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
	api.InitGrpcServer()
}
