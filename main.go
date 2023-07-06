package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"trojan-panel-core/api"
	"trojan-panel-core/app"
	"trojan-panel-core/core"
	"trojan-panel-core/dao"
	"trojan-panel-core/dao/redis"
	"trojan-panel-core/middleware"
	"trojan-panel-core/router"
)

func main() {
	serverConfig := core.Config.ServerConfig
	r := gin.Default()
	router.Router(r)
	_ = r.Run(fmt.Sprintf(":%d", serverConfig.Port))
	defer closeResource()
}

func init() {
	core.InitConfig()
	middleware.InitLog()
	dao.InitMySQL()
	dao.InitSqlLite()
	redis.InitRedis()
	middleware.InitCron()
	middleware.InitRateLimiter()
	api.InitValidator()
	api.InitGrpcServer()
	app.InitApp()
}
func closeResource() {
	dao.CloseDb()
	dao.CloseSqliteDb()
	redis.CloseRedis()
}
