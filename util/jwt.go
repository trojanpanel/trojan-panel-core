package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	redisgo "github.com/gomodule/redigo/redis"
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

func GetJWTKey() ([]byte, error) {
	get := redis.Client.String.
		Get("trojan-panel:jwt-key")
	reply, err := get.Bytes()
	if err != nil && err != redisgo.ErrNil {
		return nil, errors.New(constant.SysError)
	}
	if len(reply) > 0 {
		return reply, nil
	}
	return nil, errors.New(constant.SysError)
}
