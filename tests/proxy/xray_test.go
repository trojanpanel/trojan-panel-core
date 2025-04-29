package proxy

import (
	"fmt"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/proxy/trojan"
	"testing"
	"trojan-core/proxy"
)

func TestDownloadXray(t *testing.T) {
	if err := proxy.DownloadXray(""); err != nil {
		println(err.Error())
	}
}

func TestXrayQueryStats(t *testing.T) {
	xrayApi := proxy.NewXrayApi("18080")
	stats, err := xrayApi.QueryStats("", false)
	if err != nil {
		return
	}
	for _, stat := range stats {
		println(fmt.Sprintf("%s -> %d", stat.Name, stat.Value))
	}
}

func TestXrayDeleteUser(t *testing.T) {
	password := "love@example.com"
	xrayApi := proxy.NewXrayApi("18080")
	if err := xrayApi.DeleteUser(password); err != nil {
		return
	}
}

func TestXrayAddUser(t *testing.T) {
	password := "123456"
	xrayApi := proxy.NewXrayApi("18080")
	if err := xrayApi.AddUser(password,
		serial.ToTypedMessage(&trojan.Account{
			Password: password,
		}),
	); err != nil {
		return
	}
}
