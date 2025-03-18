package service

import (
	"encoding/base64"
	"trojan-core/proxy"
	"trojan-core/util"
)

func HandleHysteriaAccountAuth(auth string) (bool, error) {
	authUsers, err := ListAuthUsers()
	if err != nil {
		return false, err
	}
	return util.ArrContain(authUsers, base64.StdEncoding.EncodeToString([]byte(auth))), nil
}

func handleHysteriaAccountTraffic(apiPort string) {
	users, err := proxy.NewHysteriaApi(apiPort).ListUsers(true, "")
	if err != nil {
		return
	}
	for user, traffic := range users {
		go XAddAccountTraffic(base64.StdEncoding.EncodeToString([]byte(user)), traffic.Tx, traffic.Rx)
	}
}
