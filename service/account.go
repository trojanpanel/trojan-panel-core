package service

import "trojan-core/proxy"

func HandleAccount() {
	proxy.XrayCmdMap.Range(func(key, value any) bool {
		go handleXrayAccountAuth(key.(string))
		go handleXrayAccountTraffic(key.(string))
		return true
	})
	proxy.HysteriaCmdMap.Range(func(key, value any) bool {
		go handleHysteriaAccountAuth(key.(string))
		go handleHysteriaAccountTraffic(key.(string))
		return true
	})
	proxy.NaiveProxyCmdMap.Range(func(key, value any) bool {
		go handleNaiveProxyAccountAuth(key.(string))
		go handleNaiveProxyAccountTraffic(key.(string))
		return true
	})
}
