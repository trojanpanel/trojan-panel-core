package api

import (
	"github.com/gin-gonic/gin"
	"trojan-panel-core/model/constant"
	"trojan-panel-core/model/dto"
	"trojan-panel-core/model/vo"
	"trojan-panel-core/service"
)

func Hysteria2Api(c *gin.Context) {
	var hysteria2AuthDto dto.Hysteria2AuthDto
	_ = c.ShouldBindJSON(&hysteria2AuthDto)
	if err := validate.Struct(&hysteria2AuthDto); err != nil {
		vo.Hysteria2ApiFail(constant.ValidateFailed, c)
		return
	}
	//base64DecodeStr, err := base64.StdEncoding.DecodeString()
	//if err != nil {
	//	vo.Hysteria2ApiFail(constant.ValidateFailed, c)
	//	return
	//}
	//pass := string(base64DecodeStr)
	accountHysteria2Vo, err := service.SelectAccountByPass(*hysteria2AuthDto.Auth)
	if err != nil || accountHysteria2Vo == nil {
		vo.Hysteria2ApiFail("", c)
		return
	}
	vo.Hysteria2ApiSuccess(*hysteria2AuthDto.Auth, c)
}
