package service

import "trojan-panel-core/util"

func GetServerStats() (cpuUsed float64, memUsed float64, diskUsed float64, err error) {
	cpuUsed, err = util.GetCpuPercent()
	memUsed, err = util.GetMemPercent()
	diskUsed, err = util.GetDiskPercent()
	return
}
