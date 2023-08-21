package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var limit *limiter.Limiter

func RateLimiterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(limit, c.Writer, c.Request)
		if httpError != nil {
			logrus.Warnf("request too fast ip: %s", c.ClientIP())
			c.Abort()
			return
		}
		c.Next()
	}
}

func InitRateLimiter() {
	limit = tollbooth.NewLimiter(5, nil)
}
