package util

import (
	"testing"
	"trojan-core/util"
)

func TestGenerateUUID(t *testing.T) {
	println(util.GenerateUUID("love@example.com"))
}
