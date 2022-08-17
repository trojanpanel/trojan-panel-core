package dto

type HysteriaConfigDto struct {
	ApiPort  uint // api端口
	Port     uint // hysteria端口
	Protocol string
	Ip       string
	UpMbps   int
	DownMbps int
}

// HysteriaAutoDto hysteria api端口
type HysteriaAutoDto struct {
	Payload *string `json:"payload" validate:"required"`
}
