package util

import (
	"testing"
	"trojan-core/util"
)

func TestGetFileNameWithoutExt(t *testing.T) {
	print(util.GetFileNameWithoutExt("test\\1.go"))
}
