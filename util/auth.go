package util

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"trojan-core/model/constant"
)

func AuthRequest(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf(constant.UnauthorizedError)
	}
	var token string
	if val, ok := md["token"]; ok {
		token = val[0]
	}
	myClaims, err := ParseToken(token)
	if err != nil {
		return fmt.Errorf(constant.UnauthorizedError)
	}
	if myClaims.AccountVo.Deleted == 1 || !ArrContain(myClaims.AccountVo.Roles, "admin") {
		return fmt.Errorf(constant.ForbiddenError)
	}
	return nil
}
