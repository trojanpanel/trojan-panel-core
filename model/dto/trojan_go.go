package dto

type TrojanGoConfigDto struct {
	Port            uint
	Domain          string
	Sni             string
	MuxEnable       uint
	WebsocketEnable uint
	WebsocketPath   string
	WebsocketHost   string
	SSEnable        uint
	SSMethod        string
	SSPassword      string
	ApiPort         uint
}

type TrojanGoAddUserDto struct {
	Hash               string
	IpLimit            int
	UploadTraffic      int
	DownloadTraffic    int
	UploadSpeedLimit   int
	DownloadSpeedLimit int
}
