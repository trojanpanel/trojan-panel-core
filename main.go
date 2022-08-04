package main

import (
	"github.com/gin-gonic/gin"
	"xray-manage/core"
	"xray-manage/core/xray"
	"xray-manage/dao"
	"xray-manage/middleware"
	"xray-manage/util"
)

func main() {
	r := gin.Default()
	_ = r.Run(":8082")
}

func init() {
	util.InitFile()
	core.InitConfig()
	middleware.InitLog()
	dao.InitMySQL()
	xray.InitGrpcClientConn()
	middleware.InitCron()
}
