package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	redisgo "github.com/gomodule/redigo/redis"
	"trojan-core/dao"
	"trojan-core/model/constant"
	"trojan-core/model/vo"
)

type MyClaims struct {
	AccountVo vo.AccountVo `json:"accountVo"`
	jwt.StandardClaims
}

func ParseToken(tokenString string) (*MyClaims, error) {
	mySecret, err := GetJWTKey()
	if err != nil {
		return nil, errors.New(constant.SysError)
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, errors.New(constant.IllegalTokenError)
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New(constant.TokenExpiredError)
}

func GetJWTKey() ([]byte, error) {
	reply, err := dao.RedisClient.Get(context.Background(), constant.JwtKey).Bytes()
	if err != nil && err != redisgo.ErrNil {
		return nil, errors.New(constant.SysError)
	}
	if len(reply) > 0 {
		return reply, nil
	}
	return nil, errors.New(constant.SysError)
}
