package proxy

import (
	"testing"
	"trojan-core/proxy"
)

func TestDownloadHysteria(t *testing.T) {
	if err := proxy.DownloadHysteria(""); err != nil {
		println(err.Error())
	}
}
