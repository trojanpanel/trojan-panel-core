package controller

import (
	"github.com/gin-gonic/gin"
	"trojan-core/model/constant"
	"trojan-core/model/dto"
	"trojan-core/model/vo"
	"trojan-core/service"
)

func HysteriaApi(c *gin.Context) {
	var hysteria2AuthDto dto.Hysteria2AuthDto
	_ = c.ShouldBindJSON(&hysteria2AuthDto)
	if err := validate.Struct(&hysteria2AuthDto); err != nil {
		vo.HysteriaApiFail(constant.InvalidError, c)
		return
	}
	accountHysteria2Vo, err := service.SelectAccountByPass(*hysteria2AuthDto.Auth)
	if err != nil || accountHysteria2Vo == nil {
		vo.HysteriaApiFail("", c)
		return
	}
	vo.HysteriaApiSuccess(*hysteria2AuthDto.Auth, c)
}
