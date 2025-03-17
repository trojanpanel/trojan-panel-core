package controller

import (
	"github.com/gin-gonic/gin"
	"trojan-core/model/constant"
	"trojan-core/model/dto"
	"trojan-core/model/vo"
	"trojan-core/service"
)

func HysteriaApi(c *gin.Context) {
	var hysteriaAuthDto dto.HysteriaAuthDto
	_ = c.ShouldBindJSON(&hysteriaAuthDto)
	if err := validate.Struct(&hysteriaAuthDto); err != nil {
		vo.HysteriaApiFail(constant.InvalidError, c)
		return
	}
	// hysteria 认证
	auth, err := service.HandleHysteriaAccountAuth(*hysteriaAuthDto.Auth)
	if err != nil || !auth {
		vo.HysteriaApiFail("", c)
		return
	}
	vo.HysteriaApiSuccess(*hysteriaAuthDto.Auth, c)
}
