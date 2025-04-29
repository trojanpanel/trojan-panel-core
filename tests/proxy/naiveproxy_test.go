package proxy

import (
	"fmt"
	"testing"
	"trojan-core/proxy"
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
