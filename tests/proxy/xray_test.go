package proxy

import (
	"fmt"
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
	stats, err := xrayApi.QueryStats("", true)
	if err != nil {
		return
	}
	for _, stat := range stats {
		println(fmt.Sprintf("%s -> %d", stat.Name, stat.Value))
	}
}

func TestXrayDeleteUser(t *testing.T) {
	xrayApi := proxy.NewXrayApi("18080")
	if err := xrayApi.DeleteUser("love@example.com"); err != nil {
		return
	}
}
