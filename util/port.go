package util

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/common/net"
	"math/rand"
	"time"
	"trojan-panel-core/module/constant"
)

// GetPortAvailBetween 获取10000到10100之间的随机整数
func GetPortAvailBetween() (uint, error) {
	rand.Seed(time.Now().Unix())
	port := 10000
	for !IsPortAvailable(port) {
		port = port + 1
		if port > 10100 {
			return 0, errors.New(constant.PortExceed)
		}
	}
	return uint(port), nil
}

// IsPortAvailable 判断端口是否可用
func IsPortAvailable(port int) bool {
	address := fmt.Sprintf("127.0.0.1:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Errorf("port %s is taken: %s \n", address, err)
		return false
	}

	defer listener.Close()
	return true
}
