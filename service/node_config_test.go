package service

import (
	"fmt"
	"testing"
)

func TestSelectNodeConfigByNodeTypeIdAndApiPort(t *testing.T) {
	nodeConfig, err := SelectNodeConfigByNodeTypeIdAndApiPort(443, 1)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Println(nodeConfig)
}
