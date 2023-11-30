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
	Network         string      `json:"network,omitempty"`
	Security        string      `json:"security,omitempty"`
	TlsSettings     TlsSettings `json:"tlsSettings,omitempty"`
	RealitySettings TypeMessage `json:"realitySettings,omitempty"`
	WsSettings      TypeMessage `json:"wsSettings,omitempty"`
}

type TlsSettings struct {
	Certificates  []Certificate `json:"certificates,omitempty"`
	ServerName    TypeMessage   `json:"serverName,omitempty"`
	Alpn          TypeMessage   `json:"alpn,omitempty"`
	AllowInsecure TypeMessage   `json:"allowInsecure,omitempty"`
	Fingerprint   TypeMessage   `json:"fingerprint,omitempty"`
}

type Certificate struct {
	CertificateFile string `json:"certificateFile,omitempty"`
	KeyFile         string `json:"keyFile,omitempty"`
}
