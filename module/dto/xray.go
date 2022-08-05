package dto

type AddUserDto struct {
	Tag            string
	Email          string
	SSPassword     string // ss
	SSMethod       string // ss
	TrojanPassword string // trojan
	VId            string // vless & vmess
}

type AddBoundDto struct {
	Tag          string
	Port         int
	ProtocolName string
}
