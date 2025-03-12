package util

import (
	"testing"
	"trojan-core/util"
)

func TestGetFileNameWithoutExt(t *testing.T) {
	names, _ := util.ListFileNames("../api", ".go")
	for _, item := range names {
		println(item)
	}
}
