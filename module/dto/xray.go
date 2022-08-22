package dto

type XrayConfigDto struct {
	ApiPort        uint
	Port           uint
	Protocol       string
	Settings       string
	StreamSettings string
	Tag            string
	Sniffing       string
	Allocate       string
}

type XrayAddUserDto struct {
	Protocol string // 协议
	Password string
}
