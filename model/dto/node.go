package dto

type NodeAddDto struct {
	Proxy  string `json:"proxy" form:"proxy" validate:"required"`
	Config string `json:"config" form:"config" validate:"required"`
}
