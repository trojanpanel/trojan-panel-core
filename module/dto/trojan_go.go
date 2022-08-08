package dto

type TrojanGoConfigDto struct {
	Id              int // 节点id
	Ip              string
	Port            string
	ApiPort         string
	Sni             string
	MuxEnable       string
	WebsocketEnable string
	WebsocketPath   string
	WebsocketHost   string
	SSEnable        string
	SSMethod        string
	SSPassword      string
}
