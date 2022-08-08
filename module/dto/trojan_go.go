package dto

type TrojanGoConfigDto struct {
	LocalPort           string
	Sni                 string
	Domain              string
	MuxEnable           string
	WebsocketEnable     string
	WebsocketPath       string
	ShadowsocksEnable   string
	ShadowsocksMethod   string
	ShadowsocksPassword string
	ApiPort             string
}
