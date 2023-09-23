package dto

type Hysteria2ConfigDto struct {
	ApiPort     uint
	Port        uint // hysteria2 port
	Domain      string
	Obfs        string
	UpMbps      int
	DownMbps    int
	TrafficPort int
}

type Hysteria2AuthDto struct {
	Payload *string `json:"payload" validate:"required"`
}
