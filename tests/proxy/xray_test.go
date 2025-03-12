package proxy

import (
	"testing"
	"trojan-core/proxy"
)

func TestDownloadXray(t *testing.T) {
	if err := proxy.DownloadXray(""); err != nil {
		println(err.Error())
	}
}
