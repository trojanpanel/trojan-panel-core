package vo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// response object
type hysteria2Result struct {
	Ok bool   `json:"ok"`
	Id string `json:"id"`
}

func Hysteria2ApiSuccess(id string, c *gin.Context) {
	c.JSON(http.StatusOK, hysteria2Result{
		Ok: true,
		Id: id,
	})
}

func Hysteria2ApiFail(id string, c *gin.Context) {
	c.JSON(http.StatusOK, hysteria2Result{
		Ok: false,
		Id: id,
	})
}
