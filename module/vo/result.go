package vo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trojan-panel-core/module/constant"
)

// 返回的对象
type result struct {
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	TypeSuccess = "success"
	TypeError   = "error"
	TypeWarning = "warning"
)

// Success 封装成功返回的对象
func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, result{
		Code: constant.CodeSuccess,
		Type: TypeSuccess,
		Data: data,
	})
}

// Fail 封装失败返回的对象
func Fail(message string, c *gin.Context) {
	var code int
	if constant.IllegalTokenError == message {
		code = constant.CodeIllegalTokenError
	} else if constant.TokenExpiredError == message {
		code = constant.CodeTokenExpiredError
	} else if constant.UnauthorizedError == message {
		code = constant.CodeUnauthorizedError
	} else if constant.ForbiddenError == message {
		code = constant.CodeForbiddenError
	} else {
		code = constant.CodeSysError
	}
	c.JSON(http.StatusOK, result{
		Code:    code,
		Type:    TypeError,
		Message: message,
		Data:    nil,
	})
}

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
