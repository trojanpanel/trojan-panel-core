package proxy

import (
	"fmt"
	"testing"
	"trojan-core/proxy"
	"trojan-core/util"
)

func TestDownloadNaiveProxy(t *testing.T) {
	if err := proxy.DownloadNaiveProxy(""); err != nil {
		println(err.Error())
	}
}

func TestNaiveProxyListUsers(t *testing.T) {
	users, err := proxy.NewNaiveProxyApi("9090").ListUsers()
	if err != nil {
		return
	}
	fmt.Printf("user: %v", users)
}

func TestNaiveProxyHandleUser(t *testing.T) {
	authCredential := "123123:123123"
	authCredential = util.Base64Encode2(authCredential)
	if err := proxy.NewNaiveProxyApi("9090").HandleUser(authCredential, true); err != nil {
		return
	}
}
