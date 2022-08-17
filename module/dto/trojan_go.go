package dto

type TrojanGoConfigDto struct {
	Port            uint
	Ip              string
	Sni             string
	MuxEnable       string
	WebsocketEnable string
	WebsocketPath   string
	WebsocketHost   string
	SSEnable        string
	SSMethod        string
	SSPassword      string
	ApiPort         uint
}

type TrojanGoAddUserDto struct {
	Password           string // 无需加密
	IpLimit            int
	UploadTraffic      int
	DownloadTraffic    int
	UploadSpeedLimit   int
	DownloadSpeedLimit int
}
