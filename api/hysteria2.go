package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/service"
)

func Hysteria2Api(c *gin.Context) {
	var hysteria2AuthDto dto.Hysteria2AuthDto
	_ = c.ShouldBindJSON(&hysteria2AuthDto)
	if err := validate.Struct(&hysteria2AuthDto); err != nil {
		vo.Hysteria2ApiFail(constant.ValidateFailed, c)
		return
	}
	base64DecodeStr, err := base64.StdEncoding.DecodeString(*hysteria2AuthDto.Payload)
	if err != nil {
		vo.Hysteria2ApiFail(constant.ValidateFailed, c)
		return
	}
	pass := string(base64DecodeStr)
	accountHysteria2Vo, err := service.SelectAccountByPass(pass)
	if err != nil || accountHysteria2Vo == nil {
		vo.Hysteria2ApiFail(constant.UsernameOrPassError, c)
		return
	}
	vo.Hysteria2ApiSuccess("success", c)
}
