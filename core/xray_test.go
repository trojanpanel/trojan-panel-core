package core

import (
	"fmt"
	"testing"
)

func init() {
	InitGrpcClientConn()
}

func TestQueryStats(t *testing.T) {
	stats := QueryStats("", false)
	fmt.Println(stats)
}

func TestGetUserStats(t *testing.T) {
	stats := GetUserStats("123123", "uplink", false)
	fmt.Println(stats)
}

func TestGetBoundStats(t *testing.T) {
	stats := GetBoundStats("inbound", "api", "uplink", false)
	fmt.Println(stats)
}
