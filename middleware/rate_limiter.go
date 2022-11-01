package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var limit *limiter.Limiter

// RateLimiterHandler 限流中间件
func RateLimiterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 限流
		httpError := tollbooth.LimitByRequest(limit, c.Writer, c.Request)
		if httpError != nil {
			logrus.Warnf("请求太快了 ip: %s", c.ClientIP())
			c.Abort()
			return
		}
		c.Next()
	}
}

// InitRateLimiter 限流初始化
func InitRateLimiter() {
	limit = tollbooth.NewLimiter(5, nil)
}
