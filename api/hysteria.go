package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"strings"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
)

func HysteriaApi(c *gin.Context) {
	var hysteriaAutoDto dto.HysteriaAutoDto
	_ = c.ShouldBindJSON(&hysteriaAutoDto)
	if err := validate.Struct(&hysteriaAutoDto); err != nil {
		vo.HysteriaApiFail(constant.ValidateFailed, c)
		return
	}
	decodeString, err := base64.StdEncoding.DecodeString(*hysteriaAutoDto.Payload)
	if err != nil {
		vo.HysteriaApiFail(constant.ValidateFailed, c)
		return
	}
	usernameAndPass := strings.Split(string(decodeString), "&")
	usersVo, err := service.SelectUserByUsernameAndPass(&usernameAndPass[0], &usernameAndPass[1])
	if err != nil {
		vo.HysteriaApiFail(err.Error(), c)
		return
	}
	if usersVo.Deleted == 1 {
		vo.HysteriaApiFail(constant.UserDisabled, c)
		return
	}
	if usersVo != nil {
		vo.HysteriaApiSuccess(usersVo.Username, c)
		return
	}
	vo.HysteriaApiFail(constant.UsernameOrPassError, c)
}
