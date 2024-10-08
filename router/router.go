package router

import (
	"github.com/gin-gonic/gin"
	"trojan-core/api"
)

func Router(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
		// Hysteria api
		auth.POST("/hysteria", api.HysteriaApi)
		// Hysteria2 api
		auth.POST("/hysteria2", api.Hysteria2Api)
	}
}
