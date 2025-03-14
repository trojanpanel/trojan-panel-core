package service

import "trojan-core/proxy"

func HandleAccount() {
	go proxy.XrayCmdMap.Range(func(key, value any) bool {
		go handleXrayAccountTraffic(key.(string))
		go handleXrayAccountAuth(key.(string))
		return true
	})
	go proxy.HysteriaCmdMap.Range(func(key, value any) bool {
		go handleHysteriaAccountTraffic(key.(string))
		go handleHysteriaAccountAuth(key.(string))
		return true
	})
	go proxy.NaiveProxyCmdMap.Range(func(key, value any) bool {
		go handleNaiveProxyAccountTraffic(key.(string))
		go handleNaiveProxyAccountAuth(key.(string))
		return true
	})
}
