package dto

type HysteriaAuthDto struct {
	Auth *string `json:"auth" validate:"required"`
}
