package dto

type Hysteria2AuthDto struct {
	Auth *string `json:"auth" validate:"required"`
}
