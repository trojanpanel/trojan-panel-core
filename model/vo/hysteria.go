package vo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountHysteriaVo struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

// response object
type hysteriaResult struct {
	Ok bool   `json:"ok"`
	Id string `json:"id"`
}

func HysteriaApiSuccess(id string, c *gin.Context) {
	c.JSON(http.StatusOK, hysteriaResult{
		Ok: true,
		Id: id,
	})
}

func HysteriaApiFail(id string, c *gin.Context) {
	c.JSON(http.StatusOK, hysteriaResult{
		Ok: false,
		Id: id,
	})
}
