package bo

type Hysteria2UserTraffic struct {
	Tx int64 `json:"tx"` // upload
	Rx int64 `json:"rx"` // download
}

type Hysteria2User struct {
	Pass string `json:"Pass"`
	Tx   int64  `json:"tx"` // upload
	Rx   int64  `json:"rx"` // download
}

type Hysteria2Config struct {
	Listen       string               `json:"listen"`
	Tls          any                  `json:"tls"`
	Obfs         *Hysteria2ConfigObfs `json:"obfs"`
	Bandwidth    any                  `json:"bandwidth"`
	Auth         any                  `json:"auth"`
	TrafficStats any                  `json:"trafficStats"`
}

type Hysteria2ConfigObfs struct {
	Type       string     `json:"type"`
	Salamander Salamander `json:"salamander"`
}

type Salamander struct {
	Password string `json:"password"`
}
