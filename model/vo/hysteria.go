package vo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// response object
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
