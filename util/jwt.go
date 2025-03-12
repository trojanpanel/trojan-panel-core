package util

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"trojan-core/dao"
	"trojan-core/model/constant"
	"trojan-core/model/vo"
)

type MyClaims struct {
	AccountVo vo.AccountVo
	jwt.StandardClaims
}

func ParseToken(tokenString string) (*MyClaims, error) {
	mySecret, err := GetJWTKey()
	if err != nil {
		return nil, fmt.Errorf(constant.SysError)
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf(constant.IllegalTokenError)
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf(constant.InvalidError)
}

func GetJWTKey() ([]byte, error) {
	reply, err := dao.RedisClient.Get(context.Background(), constant.TokenSecret).Bytes()
	if err != nil {
		return nil, fmt.Errorf(constant.SysError)
	}
	if len(reply) > 0 {
		return reply, nil
	}
	return nil, fmt.Errorf(constant.SysError)
}
