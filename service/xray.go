package service

import (
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"regexp"
	"trojan-core/proxy"
	"trojan-core/util"
)

var userLinkRegex = regexp.MustCompile("user>>>([^>]+)>>>traffic>>>(downlink|uplink)")

func handleXrayAccount(apiPort string) {
	xrayApi := proxy.NewXrayApi(apiPort)
	stats, err := xrayApi.QueryStats("", true)
	if err != nil {
		return
	}
	var users []string
	for _, stat := range stats {
		subMatch := userLinkRegex.FindStringSubmatch(stat.Name)
		if len(subMatch) == 3 {
			users = append(users, subMatch[1])
			user := base64.StdEncoding.EncodeToString([]byte(subMatch[1]))
			isDown := subMatch[2] == "downlink"
			if isDown {
				go XAddAccountTraffic(user, 0, stat.Value)
			} else {
				go XAddAccountTraffic(user, stat.Value, 0)
			}
		}
	}
	authUsers, err := ListAuthUsers()
	if err != nil {
		return
	}
	// 删除
	deleteUsers := util.Subtract(authUsers, users)
	for _, user := range deleteUsers {
		if err = xrayApi.DeleteUser(user); err != nil {
			logrus.Errorf("xray DeleteUser err: %v", err)
		}
	}
}
