package router

import (
	"github.com/gin-gonic/gin"
	"trojan-panel-core/controller"
)

func Router(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
		auth.POST("/hysteria", controller.HysteriaApi)
	}
}
