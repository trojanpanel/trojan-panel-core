package dto

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
	Port         int
	ProtocolName string
}

type XrayConfigDto struct {
	ApiPort string // api端口
}
