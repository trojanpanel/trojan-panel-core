package hysteria2

import (
	"fmt"
	"testing"
	"trojan-panel-core/app/hysteria2"
)

var apiPort uint = 38089

func TestListUsers(t *testing.T) {
	hysteria2Api := hysteria2.NewHysteria2Api(apiPort)
	users, err := hysteria2Api.ListUsers(false)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	for index, item := range users {
		fmt.Printf("pass: %s upload: %d downlaod:%d", index, item.Tx, item.Rx)
	}
}

func TestGetUser(t *testing.T) {
	hysteria2Api := hysteria2.NewHysteria2Api(apiPort)
	user, err := hysteria2Api.GetUser("", false)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	if user != nil {
		fmt.Println("user not fount")
		return
	}
	fmt.Printf("pass: %s upload: %d downlaod:%d", user.Pass, user.Tx, user.Rx)
}
