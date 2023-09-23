package vo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// response object
type hysteria2Result struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
}

func Hysteria2ApiSuccess(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, hysteria2Result{
		Ok:  true,
		Msg: msg,
	})
}

func Hysteria2ApiFail(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, hysteria2Result{
		Ok:  false,
		Msg: msg,
	})
}
