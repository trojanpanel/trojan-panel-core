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
	if len(usernameAndPass) != 2 || len(usernameAndPass[0]) == 0 || len(usernameAndPass[1]) == 0 {
		vo.HysteriaApiFail(err.Error(), c)
		return
	}
	accountHysteriaVo, err := service.SelectAccountByUsernameAndPass(usernameAndPass[0], usernameAndPass[1])
	if err != nil || accountHysteriaVo == nil {
		vo.HysteriaApiFail(constant.UsernameOrPassError, c)
		return
	}
	vo.HysteriaApiSuccess(accountHysteriaVo.Username, c)
}
