package dto

type HysteriaConfigDto struct {
	ApiPort uint // api端口
}

// HysteriaAutoDto hysteria api端口
type HysteriaAutoDto struct {
	Payload *string `json:"payload" validate:"required"`
}
