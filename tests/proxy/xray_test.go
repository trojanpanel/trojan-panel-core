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
	stats, err := xrayApi.QueryStats("", false)
	if err != nil {
		return
	}
	for _, stat := range stats {
		println(fmt.Sprintf("%s -> %d", stat.Name, stat.Value))
	}
}
