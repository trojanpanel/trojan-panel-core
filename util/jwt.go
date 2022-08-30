package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/vo"
)

var MySecret = []byte("4eb01fa4acef754ad4fa94f4467fd343")

type MyClaims struct {
	AccountVo vo.AccountVo `json:"accountVo"`
	jwt.StandardClaims
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, errors.New(constant.IllegalTokenError)
	}
	// 校验Token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New(constant.TokenExpiredError)
}
