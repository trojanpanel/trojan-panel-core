package service

import (
	"github.com/sirupsen/logrus"
	"trojan-core/proxy"
	"trojan-core/util"
)

func handleNaiveProxyAccountAuth(apiPort string) {
	naiveProxyApi := proxy.NewNaiveProxyApi(apiPort)
	users, err := naiveProxyApi.ListUsers()
	if err != nil {
		return
	}
	authUsers, err := ListAuthUsers()
	if err != nil {
		return
	}
	addUsers := util.Subtract(users, authUsers)
	if err = naiveProxyApi.HandleUser(addUsers, true); err != nil {
		logrus.Errorf("naiveProxy addUser err: %v", err)
	}
	deleteUsers := util.Subtract(authUsers, users)
	if err = naiveProxyApi.HandleUser(deleteUsers, false); err != nil {
		logrus.Errorf("naiveProxy deleteUser err: %v", err)
	}
}
