package util

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"time"
	"trojan-panel-core/model/constant"
)

// IsPortAvailable determine whether the port is available
func IsPortAvailable(port uint, network string) bool {
	if network == "tcp" {
		listener, err := net.ListenTCP(network, &net.TCPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: int(port),
		})
		defer func() {
			if listener != nil {
				listener.Close()
			}
		}()
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
		defer func() {
			if listener != nil {
				listener.Close()
			}
		}()
		if err != nil {
			logrus.Warnf("port %d is taken err: %s", port, err)
			return false
		}
	}
	return true
}

// GetLocalIP get local IP address
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// check the ip address to determine whether the loopback address
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New(constant.GetLocalIPError)
}

// GetCpuPercent get CPU usage
func GetCpuPercent() (float64, error) {
	var err error
	percent, err := cpu.Percent(time.Second, false)
	value, err := strconv.ParseFloat(fmt.Sprintf("%.1f", percent[0]), 64)
	return value, err
}

// GetMemPercent get memory usage
func GetMemPercent() (float64, error) {
	var err error
	memInfo, err := mem.VirtualMemory()
	value, err := strconv.ParseFloat(fmt.Sprintf("%.1f", memInfo.UsedPercent), 64)
	return value, err
}

// GetDiskPercent get disk usage
func GetDiskPercent() (float64, error) {
	var err error
	parts, err := disk.Partitions(true)
	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	value, err := strconv.ParseFloat(fmt.Sprintf("%.1f", diskInfo.UsedPercent), 64)
	return value, err
}
