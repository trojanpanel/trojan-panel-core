package dto

type XrayConfigDto struct {
	Log       TypeMessage `json:"log"`
	API       TypeMessage `json:"api"`
	DNS       TypeMessage `json:"dns"`
	Routing   TypeMessage `json:"routing"`
	Policy    TypeMessage `json:"policy"`
	Inbounds  []Inbound   `json:"inbounds"`
	Outbounds TypeMessage `json:"outbounds"`
	Transport TypeMessage `json:"transport"`
	Stats     TypeMessage `json:"stats"`
	Reverse   TypeMessage `json:"reverse"`
	FakeDNS   TypeMessage `json:"fakeDns"`
}

type Inbound struct {
	Listen         TypeMessage `json:"listen"`
	Port           uint        `json:"port"`
	Protocol       string      `json:"protocol"`
	Settings       TypeMessage `json:"settings"`
	StreamSettings TypeMessage `json:"streamSettings"`
	Tag            string      `json:"tag"`
	Sniffing       TypeMessage `json:"sniffing"`
	Allocate       TypeMessage `json:"allocate"`
}

type XrayAddUserDto struct {
	Protocol       string // 协议
	Email          string // 唯一标识
	SSMethod       string // ss method
	SSPassword     string // ss password
	TrojanPassword string // trojan
	VlessId        string // vless
	VmessId        string // vmess
	VmessAlterId   string // vmess alter id
}

type XrayAddBoundDto struct {
	Tag          string
	Port         uint
	ProtocolName string
}
