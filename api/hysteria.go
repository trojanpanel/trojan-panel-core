package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
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
	base64DecodeStr, err := base64.StdEncoding.DecodeString(*hysteriaAuthDto.Payload)
	if err != nil {
		vo.HysteriaApiFail(constant.ValidateFailed, c)
		return
	}
	pass := string(base64DecodeStr)
	accountHysteriaVo, err := service.SelectAccountByPass(pass)
	if err != nil || accountHysteriaVo == nil {
		vo.HysteriaApiFail(constant.UsernameOrPassError, c)
		return
	}
	vo.HysteriaApiSuccess("success", c)
}
