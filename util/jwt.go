package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"trojan-panel-core/dao/redis"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/vo"
)

type MyClaims struct {
	AccountVo vo.AccountVo `json:"accountVo"`
	jwt.StandardClaims
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	mySecret, err := GetJWTKey()
	if err != nil {
		return nil, errors.New(constant.SysError)
	}
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
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

func GetJWTKey() (string, error) {
	get := redis.Client.String.
		Get("trojan-panel:jwt-key")
	reply, err := get.String()
	if err != nil {
		return "", errors.New(constant.SysError)
	}
	if reply != "" {
		return reply, nil
	}
	return "", errors.New(constant.SysError)
}
