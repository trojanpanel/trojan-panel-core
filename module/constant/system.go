package constant

const (
	// SaltKey 加密 盐
	SaltKey string = "well_very_funny!"

	// LogPath 日志文件夹路径
	LogPath string = "logs"

	// ConfigPath 全局配置文件夹路径
	ConfigPath string = "config"
	// ConfigFilePath 全局配置文件路径
	ConfigFilePath string = "config/config.ini"

	// SqlitePath sqlite文件夹路径
	SqlitePath string = "config/sqlite"
	// SqliteFilePath sqlite文件路径
	SqliteFilePath string = "config/sqlite/trojan_panel_core.db"

	XrayPath          string = "bin/xray/config"
	TrojanGoPath      string = "bin/trojango/config"
	HysteriaPath      string = "bin/hysteria/config"
	NaiveProxyPath    string = "bin/naiveproxy/config"
	XrayBinPath       string = "bin/xray"
	TrojanGoBinPath   string = "bin/trojango"
	HysteriaBinPath   string = "bin/hysteria"
	NaiveProxyBinPath string = "bin/naiveproxy"
	DownloadBaseUrl   string = "https://github.com/trojanpanel/install-script/releases/download/v1.0.0"

	TrojanPanelCoreVersion = "v2.0.3"
)
