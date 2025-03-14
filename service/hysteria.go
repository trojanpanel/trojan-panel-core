package service

import (
	"context"
	"trojan-core/dao"
	"trojan-core/model/constant"
	"trojan-core/proxy"
	"trojan-core/util"
)

func handleHysteriaAccountAuth(apiPort string) {
	users, err := proxy.NewHysteriaApi(apiPort).OnlineUsers("")
	if err != nil {
		return
	}
	if len(users) > 0 {
		i := 0
		usernames := make([]string, len(users))
		for k := range users {
			usernames[i] = k
			i++
		}
		result, err := dao.RedisClient.LRange(context.Background(), constant.AccountAuth, 0, -1).Result()
		if err != nil {
			return
		}
		kickUsernames := util.Subtract(result, usernames)
		if err = proxy.NewHysteriaApi(apiPort).KickUsers(kickUsernames, ""); err != nil {
			return
		}
	}
}

func handleHysteriaAccountTraffic(apiPort string) {
	users, err := proxy.NewHysteriaApi(apiPort).ListUsers(true, "")
	if err != nil {
		return
	}
	if len(users) > 0 {
		for username, traffic := range users {
			go XAddAccountTraffic(username, traffic.Tx, traffic.Rx)
		}
	}
}
