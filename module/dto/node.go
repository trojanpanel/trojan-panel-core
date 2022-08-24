package dto

type NodeAddDto struct {
	NodeType uint `json:"nodeType" form:"nodeType" validate:"required"`

	XrayPort           uint   `json:"XrayPort" form:"XrayPort" validate:"omitempty"`
	XrayProtocol       string `json:"XrayProtocol" form:"XrayProtocol" validate:"omitempty"`
	XraySettings       string `json:"xraySettings" form:"xraySettings" validate:"omitempty"`
	XrayStreamSettings string `json:"xrayStreamSettings" form:"xrayStreamSettings" validate:"omitempty"`
	XrayTag            string `json:"xrayTag" form:"xrayTag" validate:"omitempty"`
	XraySniffing       string `json:"xraySniffing" form:"xraySniffing" validate:"omitempty"`
	XrayAllocate       string `json:"xrayAllocate" form:"xrayAllocate" validate:"omitempty"`

	TrojanGoPort            uint   `json:"trojanGoPort" form:"trojanGoPort" validate:"omitempty"`
	TrojanGoIp              string `json:"trojanGoIp" form:"trojanGoIp" validate:"omitempty"`
	TrojanGoSni             string `json:"trojanGoSni" form:"trojanGoSni" validate:"omitempty"`
	TrojanGoMuxEnable       string `json:"trojanGoMuxEnable" form:"trojanGoMuxEnable" validate:"omitempty"`
	TrojanGoWebsocketEnable string `json:"trojanGoWebsocketEnable" form:"trojanGoWebsocketEnable" validate:"omitempty"`
	TrojanGoWebsocketPath   string `json:"trojanGoWebsocketPath" form:"trojanGoWebsocketPath" validate:"omitempty"`
	TrojanGoWebsocketHost   string `json:"trojanGoWebsocketHost" form:"trojanGoWebsocketHost" validate:"omitempty"`
	TrojanGoSSEnable        string `json:"trojanGoSSEnable" form:"trojanGoSSEnable" validate:"omitempty"`
	TrojanGoSSMethod        string `json:"trojanGoSSMethod" form:"trojanGoSSMethod" validate:"omitempty"`
	TrojanGoSSPassword      string `json:"trojanGoSSPassword" form:"trojanGoSSPassword" validate:"omitempty"`

	HysteriaPort     uint   `json:"hysteriaPort" form:"hysteriaPort" validate:"omitempty"`
	HysteriaProtocol string `json:"hysteriaProtocol" form:"hysteriaProtocol" validate:"omitempty"`
	HysteriaIp       string `json:"hysteriaIp" form:"hysteriaIp" validate:"omitempty"`
	HysteriaUpMbps   int    `json:"hysteriaUpMbps" form:"hysteriaUpMbps" validate:"omitempty"`
	HysteriaDownMbps int    `json:"hysteriaDownMbps" form:"hysteriaDownMbps" validate:"omitempty"`
}

type NodeRemoveDto struct {
	NodeType uint `json:"nodeType" form:"nodeType" validate:"required"`
	ApiPort  uint `json:"apiPort" form:"apiPort" validate:"required"`
}
