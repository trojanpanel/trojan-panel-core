package util

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(SHA224String("sysadmin123456"))
}

func TestSha1(t *testing.T) {
	fmt.Println(Sha1String("sysadmin123456"))
	fmt.Println(Sha1Match("SC2FYsVxD52nLnQ-nEAhNvM8ou1H1MaEGe9Q1UqgPNzBxeNX", "sysadmin123456"))
}
