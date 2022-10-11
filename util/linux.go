package util

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"trojan-panel-core/module/constant"
)

// IsPortAvailable 判断端口是否可用
func IsPortAvailable(port uint, network string) bool {
	address := fmt.Sprintf("127.0.0.1:%d", port)
	listener, err := net.Listen(network, address)
	if err != nil {
		logrus.Errorf("port %s is taken: %s \n", address, err)
		return false
	}

	defer listener.Close()
	return true
}

// GetLocalIP 获取本机IP地址
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New(constant.GetLocalIPError)
}
