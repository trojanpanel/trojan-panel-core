package dto

type Hysteria2ConfigDto struct {
	ApiPort      uint
	Port         uint // hysteria2 port
	Domain       string
	ObfsPassword string
	UpMbps       int
	DownMbps     int
}

type Hysteria2AuthDto struct {
	Auth *string `json:"auth" validate:"required"`
}
