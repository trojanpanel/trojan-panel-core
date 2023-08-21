package dto

type HysteriaConfigDto struct {
	ApiPort  uint
	Port     uint // hysteria port
	Protocol string
	Domain   string
	Obfs     string
	UpMbps   int
	DownMbps int
}

type HysteriaAuthDto struct {
	Payload *string `json:"payload" validate:"required"`
}
