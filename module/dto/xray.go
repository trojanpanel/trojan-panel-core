package dto

type XrayConfigDto struct {
	ApiPort  uint
	Port     uint
	Protocol string
}

type XrayAddUserDto struct {
	Protocol string // 协议
	Password string
}
