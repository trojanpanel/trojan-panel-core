package controller

import (
	"github.com/gin-gonic/gin"
	"trojan-panel-core/model/constant"
	"trojan-panel-core/model/dto"
	"trojan-panel-core/model/vo"
)

func HysteriaApi(c *gin.Context) {
	var hysteria2AuthDto dto.Hysteria2AuthDto
	_ = c.ShouldBindJSON(&hysteria2AuthDto)
	if err := validate.Struct(&hysteria2AuthDto); err != nil {
		vo.HysteriaApiFail(constant.InvalidError, c)
		return
	}
	// hysteria 认证
	vo.HysteriaApiSuccess(*hysteria2AuthDto.Auth, c)
}
