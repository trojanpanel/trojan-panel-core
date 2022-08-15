package dto

type TrojanGoConfigDto struct {
	Id              uint // 节点id
	Ip              string
	Port            uint
	ApiPort         uint
	Sni             string
	MuxEnable       string
	WebsocketEnable string
	WebsocketPath   string
	WebsocketHost   string
	SSEnable        string
	SSMethod        string
	SSPassword      string
}

type TrojanGoAddUserDto struct {
	Password           string // 无需加密
	IpLimit            int
	UploadTraffic      int
	DownloadTraffic    int
	UploadSpeedLimit   int
	DownloadSpeedLimit int
}
