package dto

type XrayConfigDto struct {
	ApiPort uint // api端口
}

type XrayAddUserDto struct {
	Tag            string
	Email          string
	SSPassword     string // ss
	SSMethod       string // ss
	TrojanPassword string // trojan
	VId            string // vless & vmess
}

type XrayAddBoundDto struct {
	Tag          string
	Port         uint
	ProtocolName string
}
