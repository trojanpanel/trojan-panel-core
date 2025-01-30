package controller

import (
	"github.com/gin-gonic/gin"
	"trojan-panel-core/model/constant"
	"trojan-panel-core/model/dto"
	"trojan-panel-core/model/vo"
)

func HysteriaApi(c *gin.Context) {
	var hysteriaAuthDto dto.HysteriaAuthDto
	_ = c.ShouldBindJSON(&hysteriaAuthDto)
	if err := validate.Struct(&hysteriaAuthDto); err != nil {
		vo.HysteriaApiFail(constant.InvalidError, c)
		return
	}
	// hysteria 认证
	vo.HysteriaApiSuccess(*hysteriaAuthDto.Auth, c)
}
