package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"regexp"
	"trojan-core/dao"
	"trojan-core/model/constant"
	"trojan-core/proxy"
	"trojan-core/util"
)

var userLinkRegex = regexp.MustCompile("user>>>([^>]+)>>>traffic>>>(downlink|uplink)")

func handleXrayAccountAuth(apiPort string) {
	xrayApi := proxy.NewXrayApi(apiPort)
	stats, err := xrayApi.QueryStats("", false)
	if err != nil {
		return
	}
	var auths []string
	for _, stat := range stats {
		subMatch := userLinkRegex.FindStringSubmatch(stat.Name)
		if len(subMatch) == 3 {
			auths = append(auths, util.SHA224String(subMatch[1]))
		}
	}
	result, err := dao.RedisClient.LRange(context.Background(), constant.AccountAuth, 0, -1).Result()
	if err != nil {
		return
	}
	deleteAuths := util.Subtract(result, auths)
	for _, item := range deleteAuths {
		if err = xrayApi.DeleteUser(item); err != nil {
			logrus.Errorf("xray DeleteUser err: %v", err)
			continue
		}
	}
}

func handleXrayAccountTraffic(apiPort string) {
	xrayApi := proxy.NewXrayApi(apiPort)
	stats, err := xrayApi.QueryStats("", false)
	if err != nil {
		return
	}
	for _, stat := range stats {
		subMatch := userLinkRegex.FindStringSubmatch(stat.Name)
		if len(subMatch) == 3 {
			username := util.SHA224String(subMatch[1])
			isDown := subMatch[2] == "downlink"
			if isDown {
				go XAddAccountTraffic(username, 0, stat.Value)
			} else {
				go XAddAccountTraffic(username, stat.Value, 0)
			}
		}
	}
}
