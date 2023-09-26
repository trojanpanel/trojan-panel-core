package util

import (
	"fmt"
	"testing"
	"trojan-panel-core/util"
)

func Test(t *testing.T) {
	fmt.Println(util.SHA224String("sysadmin123456"))
}

func TestSha1(t *testing.T) {
	fmt.Println(util.Sha1String("sysadmin123456"))
	fmt.Println(util.Sha1Match("SC2FYsVxD52nLnQ-nEAhNvM8ou1H1MaEGe9Q1UqgPNzBxeNX", "sysadmin123456"))
}
