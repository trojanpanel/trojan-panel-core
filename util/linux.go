package util

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net"
	"trojan-panel-core/module/constant"
)

// IsPortAvailable 判断端口是否可用
func IsPortAvailable(port uint, network string) bool {
	if network == "tcp" {
		listener, err := net.ListenTCP(network, &net.TCPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: int(port),
		})
		defer listener.Close()
		if err != nil {
			logrus.Warnf("port %d is taken err: %s", port, err)
			return false
		}
	}
	if network == "udp" {
		listener, err := net.ListenUDP("udp", &net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: int(port),
		})
		defer listener.Close()
		if err != nil {
			logrus.Warnf("port %d is taken err: %s", port, err)
			return false
		}
	}
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
