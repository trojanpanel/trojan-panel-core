package dto

type AddUserDto struct {
	Tag            string
	Email          string
	SSPassword     string // ss
	SSMethod       string // ss
	TrojanPassword string // trojan
	VId            string // vless & vmess
}
