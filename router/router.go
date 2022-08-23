package router

import (
	"github.com/gin-gonic/gin"
	"trojan-panel-core/api"
	"trojan-panel-core/middleware"
)

func Router(router *gin.Engine) {
	router.Use(middleware.RateLimiterHandler(), middleware.LogHandler())
	auth := router.Group("/api/auth")
	{
		// Hysteria api
		auth.POST("/hysteria", api.HysteriaApi)
	}
	router.Use(middleware.JWTHandler())
	node := router.Group("/api/node")
	{
		node.POST("/add", api.AddNode)
		node.POST("/remove", api.RemoveNode)
	}
}
