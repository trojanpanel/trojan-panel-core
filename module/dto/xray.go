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
	Template       string
}

type XrayAddUserDto struct {
	Protocol string
	Password string
}
