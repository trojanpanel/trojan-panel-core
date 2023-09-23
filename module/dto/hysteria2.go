package dto

type Hysteria2ConfigDto struct {
	ApiPort  uint
	Port     uint // hysteria2 port
	Protocol string
	Domain   string
	Obfs     string
	UpMbps   int
	DownMbps int
}

type Hysteria2AuthDto struct {
	Payload *string `json:"payload" validate:"required"`
}
