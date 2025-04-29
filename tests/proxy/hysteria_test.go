package proxy

import (
	"fmt"
	"testing"
	"trojan-core/proxy"
)

func TestDownloadHysteria(t *testing.T) {
	if err := proxy.DownloadHysteria(""); err != nil {
		println(err.Error())
	}
}

func TestHysteriaListUsers(t *testing.T) {
	users, err := proxy.NewHysteriaApi("9999").ListUsers(true)
	if err != nil {
		return
	}
	for _, user := range users {
		fmt.Printf("user: %v", user)
	}
}
