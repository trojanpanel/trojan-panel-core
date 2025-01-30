package util

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"trojan-panel-core/dao"
	"trojan-panel-core/model/constant"
	"trojan-panel-core/model/vo"
)

type MyClaims struct {
	AccountVo vo.AccountVo
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
	return nil, errors.New(constant.InvalidError)
}

func GetJWTKey() ([]byte, error) {
	reply, err := dao.RedisClient.Get(context.Background(), constant.TokenSecret).Bytes()
	if err != nil {
		return nil, errors.New(constant.SysError)
	}
	if len(reply) > 0 {
		return reply, nil
	}
	return nil, errors.New(constant.SysError)
}
