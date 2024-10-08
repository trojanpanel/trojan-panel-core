package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"trojan-core/api"
	"trojan-core/app"
	"trojan-core/core"
	"trojan-core/dao"
	"trojan-core/dao/redis"
	"trojan-core/middleware"
	"trojan-core/router"
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
