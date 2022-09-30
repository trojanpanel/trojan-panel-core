package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/metadata"
	"trojan-panel-core/dao/redis"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

var validate *validator.Validate

func InitValidator() {
	// Validate为单例对象
	validate = validator.New()
}

// token认证
func authRequest(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New(constant.UnauthorizedError)
	}
	var token string
	if val, ok := md["token"]; ok {
		token = val[0]
	}
	myClaims, err := util.ParseToken(token)
	if err != nil {
		return errors.New(constant.UnauthorizedError)
	}
	get := redis.Client.String.
		Get(fmt.Sprintf("trojan-panel:token:%s", myClaims.AccountVo.Username))
	result, err := get.String()
	if err != nil || result == "" {
		return errors.New(constant.IllegalTokenError)
	}
	return nil
}
