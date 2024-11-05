package api

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"trojan-core/model/constant"
	"trojan-core/util"
)

// Token Authentication
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
	if myClaims.AccountVo.Deleted == 1 || !util.IsAdmin(myClaims.AccountVo.Roles) {
		return errors.New(constant.ForbiddenError)
	}
	return nil
}
