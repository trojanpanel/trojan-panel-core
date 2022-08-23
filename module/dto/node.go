package dto

type NodeAddDto struct {
	NodeType uint `json:"nodeType" form:"nodeType" validate:"required"`

	XrayPort           uint   `json:"XrayPort" form:"XrayPort"`
	XrayProtocol       string `json:"XrayProtocol" form:"XrayProtocol"`
	XraySettings       string `json:"xraySettings" form:"xraySettings"`
	XrayStreamSettings string `json:"xrayStreamSettings" form:"xrayStreamSettings"`
	XrayTag            string `json:"xrayTag" form:"xrayTag"`
	XraySniffing       string `json:"xraySniffing" form:"xraySniffing"`
	XrayAllocate       string `json:"xrayAllocate" form:"xrayAllocate"`

	TrojanGoPort            uint   `json:"trojanGoPort" form:"trojanGoPort"`
	TrojanGoIp              string `json:"trojanGoIp" form:"trojanGoIp"`
	TrojanGoSni             string `json:"trojanGoSni" form:"trojanGoSni"`
	TrojanGoMuxEnable       string `json:"trojanGoMuxEnable" form:"trojanGoMuxEnable"`
	TrojanGoWebsocketEnable string `json:"trojanGoWebsocketEnable" form:"trojanGoWebsocketEnable"`
	TrojanGoWebsocketPath   string `json:"trojanGoWebsocketPath" form:"trojanGoWebsocketPath"`
	TrojanGoWebsocketHost   string `json:"trojanGoWebsocketHost" form:"trojanGoWebsocketHost"`
	TrojanGoSSEnable        string `json:"trojanGoSSEnable" form:"trojanGoSSEnable"`
	TrojanGoSSMethod        string `json:"trojanGoSSMethod" form:"trojanGoSSMethod"`
	TrojanGoSSPassword      string `json:"trojanGoSSPassword" form:"trojanGoSSPassword"`

	HysteriaPort     uint   `json:"hysteriaPort" form:"hysteriaPort"`
	HysteriaProtocol string `json:"hysteriaProtocol" form:"hysteriaProtocol"`
	HysteriaIp       string `json:"hysteriaIp" form:"hysteriaIp"`
	HysteriaUpMbps   int    `json:"hysteriaUpMbps" form:"hysteriaUpMbps"`
	HysteriaDownMbps int    `json:"hysteriaDownMbps" form:"hysteriaDownMbps"`
}

type NodeRemoveDto struct {
	NodeType uint `json:"nodeType" form:"nodeType" validate:"required"`
	ApiPort  uint `json:"apiPort" form:"apiPort" validate:"required"`
}
