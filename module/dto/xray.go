package dto

type AddUserDto struct {
	Tag            string
	Email          string
	SSPassword     string // ss
	SSMethod       string // ss
	TrojanPassword string // trojan
	Id             string // vless & vmess
}
