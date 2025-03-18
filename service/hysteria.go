package service

import (
	"context"
	"trojan-core/dao"
	"trojan-core/model/constant"
	"trojan-core/proxy"
	"trojan-core/util"
)

func HandleHysteriaAccountAuth(auth string) (bool, error) {
	result, err := dao.RedisClient.LRange(context.Background(), constant.AccountAuth, 0, -1).Result()
	if err != nil {
		return false, err
	}
	return util.ArrContain(result, util.SHA224String(auth)), nil
}

func handleHysteriaAccountTraffic(apiPort string) {
	users, err := proxy.NewHysteriaApi(apiPort).ListUsers(true, "")
	if err != nil {
		return
	}
	if len(users) > 0 {
		for user, traffic := range users {
			go XAddAccountTraffic(user, traffic.Tx, traffic.Rx)
		}
	}
}
