package constant

const (
	// LogPath log folder path
	LogPath string = "logs/"

	SqliteDBPath = "data/trojan_core.db"

	XrayPath          string = "bin/xray/config"
	TrojanGoPath      string = "bin/trojango/config"
	HysteriaPath      string = "bin/hysteria/config"
	Hysteria2Path     string = "bin/hysteria2/config"
	NaiveProxyPath    string = "bin/naiveproxy/config"
	XrayBinPath       string = "bin/xray"
	TrojanGoBinPath   string = "bin/trojango"
	HysteriaBinPath   string = "bin/hysteria"
	Hysteria2BinPath  string = "bin/hysteria2"
	NaiveProxyBinPath string = "bin/naiveproxy"

	JwtKey = "trojan-panel:jwt-key"

	Version = "v3.0.0"
)
