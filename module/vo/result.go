package vo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回的对象
type hysteriaResult struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
}

func HysteriaApiSuccess(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, hysteriaResult{
		Ok:  true,
		Msg: msg,
	})
}

func HysteriaApiFail(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, hysteriaResult{
		Ok:  false,
		Msg: msg,
	})
}
