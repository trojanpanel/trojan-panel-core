package dto

import "github.com/xtls/xray-core/proxy/shadowsocks"

type XrayConfigDto struct {
	ApiPort        uint
	Port           uint
	Protocol       string
	Settings       string
	StreamSettings string
	Tag            string
	Sniffing       string
	Allocate       string
	Template       string // 模板
}

type XrayAddUserDto struct {
	Protocol   string // 协议
	Password   string
	CipherType shadowsocks.CipherType // 加密方式
}
