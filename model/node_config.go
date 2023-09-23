package model

type NodeConfig struct {
	Id           uint   `json:"id" ddb:"id"`
	ApiPort      uint   `json:"apiPort" ddb:"api_port"`
	NodeTypeId   uint   `json:"nodeTypeId" ddb:"node_type_id"`
	Protocol     string `json:"protocol" ddb:"protocol"`
	XrayFlow     string `json:"xrayFlow" ddb:"xray_flow"`
	XraySSMethod string `json:"xraySSMethod" ddb:"xray_ss_method"`
}
