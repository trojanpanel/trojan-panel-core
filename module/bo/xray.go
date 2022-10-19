package bo

type XrayConfigBo struct {
	Log       TypeMessage `json:"log"`
	API       TypeMessage `json:"api"`
	DNS       TypeMessage `json:"dns"`
	Routing   TypeMessage `json:"routing"`
	Policy    TypeMessage `json:"policy"`
	Inbounds  []InboundBo `json:"inbounds"`
	Outbounds TypeMessage `json:"outbounds"`
	Transport TypeMessage `json:"transport"`
	Stats     TypeMessage `json:"stats"`
	Reverse   TypeMessage `json:"reverse"`
	FakeDNS   TypeMessage `json:"fakeDns"`
}

type InboundBo struct {
	Listen         string      `json:"listen"`
	Port           uint        `json:"port"`
	Protocol       string      `json:"protocol"`
	Settings       TypeMessage `json:"settings"`
	StreamSettings TypeMessage `json:"streamSettings"`
	Tag            string      `json:"tag"`
	Sniffing       TypeMessage `json:"sniffing"`
	Allocate       TypeMessage `json:"allocate"`
}

type StreamSettings struct {
	Network      string       `json:"network"`
	Security     string       `json:"security"`
	TlsSettings  TlsSettings  `json:"tlsSettings"`
	XtlsSettings XtlsSettings `json:"xtlsSettings"`
}

type TlsSettings struct {
	ServerName   string        `json:"serverName"`
	Certificates []Certificate `json:"certificates"`
}

type XtlsSettings struct {
	ServerName   string        `json:"serverName"`
	Certificates []Certificate `json:"certificates"`
}

type Certificate struct {
	CertificateFile string `json:"certificateFile"`
	KeyFile         string `json:"keyFile"`
}
