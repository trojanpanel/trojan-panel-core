package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"strings"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/service"
)

func HysteriaApi(c *gin.Context) {
	var hysteriaAuthDto dto.HysteriaAuthDto
	_ = c.ShouldBindJSON(&hysteriaAuthDto)
	if err := validate.Struct(&hysteriaAuthDto); err != nil {
		vo.HysteriaApiFail(constant.ValidateFailed, c)
		return
	}
	decodeString, err := base64.StdEncoding.DecodeString(*hysteriaAuthDto.Payload)
	if err != nil {
		vo.HysteriaApiFail(constant.ValidateFailed, c)
		return
	}
	usernameAndPass := strings.Split(string(decodeString), "&")
	accountVo, err := service.SelectAccountByUsernameAndPass(&usernameAndPass[0], &usernameAndPass[1])
	if err != nil {
		vo.HysteriaApiFail(err.Error(), c)
		return
	}
	if accountVo != nil {
		vo.HysteriaApiSuccess(accountVo.Username, c)
		return
	}
	vo.HysteriaApiFail(constant.UsernameOrPassError, c)
}
