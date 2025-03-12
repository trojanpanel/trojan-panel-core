package proxy

import (
	"testing"
	"trojan-core/proxy"
)

func TestDownloadNaiveProxy(t *testing.T) {
	if err := proxy.DownloadNaiveProxy(""); err != nil {
		println(err.Error())
	}
}
