package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/service"
	"trojan-panel-core/util"
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
	aesDecodeStr, err := util.AesDecode(string(base64DecodeStr))
	if err != nil {
		vo.HysteriaApiFail(constant.ValidateFailed, c)
		return
	}
	accountHysteriaVo, err := service.SelectAccountByUsernameAndPass(aesDecodeStr)
	if err != nil || accountHysteriaVo == nil {
		vo.HysteriaApiFail(constant.UsernameOrPassError, c)
		return
	}
	vo.HysteriaApiSuccess(accountHysteriaVo.Username, c)
}
