package dto

type HysteriaConfigDto struct {
	ApiPort  uint // api端口
	Port     uint // hysteria端口
	Protocol string
	Domain   string
	Obfs     string
	UpMbps   int
	DownMbps int
}

// HysteriaAuthDto hysteria api端口
type HysteriaAuthDto struct {
	Payload *string `json:"payload" validate:"required"`
}
