package dto

type NodeAddDto struct {
	NodeTypeId uint   `json:"nodeType" form:"nodeType" validate:"required"`
	Port       uint   `json:"port" form:"port" validate:"omitempty"`
	Domain     string `json:"domain" form:"domain" validate:"omitempty"`

	XrayProtocol       string `json:"xrayProtocol" form:"xrayProtocol" validate:"omitempty"`
	XraySettings       string `json:"xraySettings" form:"xraySettings" validate:"omitempty"`
	XrayStreamSettings string `json:"xrayStreamSettings" form:"xrayStreamSettings" validate:"omitempty"`
	XrayTag            string `json:"xrayTag" form:"xrayTag" validate:"omitempty"`
	XraySniffing       string `json:"xraySniffing" form:"xraySniffing" validate:"omitempty"`
	XrayAllocate       string `json:"xrayAllocate" form:"xrayAllocate" validate:"omitempty"`

	TrojanGoSni             string `json:"trojanGoSni" form:"trojanGoSni" validate:"omitempty"`
	TrojanGoMuxEnable       uint   `json:"trojanGoMuxEnable" form:"trojanGoMuxEnable" validate:"omitempty"`
	TrojanGoWebsocketEnable uint   `json:"trojanGoWebsocketEnable" form:"trojanGoWebsocketEnable" validate:"omitempty"`
	TrojanGoWebsocketPath   string `json:"trojanGoWebsocketPath" form:"trojanGoWebsocketPath" validate:"omitempty"`
	TrojanGoWebsocketHost   string `json:"trojanGoWebsocketHost" form:"trojanGoWebsocketHost" validate:"omitempty"`
	TrojanGoSSEnable        uint   `json:"trojanGoSSEnable" form:"trojanGoSSEnable" validate:"omitempty"`
	TrojanGoSSMethod        string `json:"trojanGoSSMethod" form:"trojanGoSSMethod" validate:"omitempty"`
	TrojanGoSSPassword      string `json:"trojanGoSSPassword" form:"trojanGoSSPassword" validate:"omitempty"`

	HysteriaProtocol string `json:"hysteriaProtocol" form:"hysteriaProtocol" validate:"omitempty"`
	HysteriaUpMbps   int    `json:"hysteriaUpMbps" form:"hysteriaUpMbps" validate:"omitempty"`
	HysteriaDownMbps int    `json:"hysteriaDownMbps" form:"hysteriaDownMbps" validate:"omitempty"`
}

type NodeRemoveDto struct {
	NodeType uint `json:"nodeType" form:"nodeType" validate:"required"`
	ApiPort  uint `json:"apiPort" form:"apiPort" validate:"required"`
}
