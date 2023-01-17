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
	// XrayTemplateFilePath Xray模板
	XrayTemplateFilePath string = "config/template-xray.json"

	XrayPath          string = "bin/xray/config"
	TrojanGoPath      string = "bin/trojango/config"
	HysteriaPath      string = "bin/hysteria/config"
	NaiveProxyPath    string = "bin/naiveproxy/config"
	XrayBinPath       string = "bin/xray"
	TrojanGoBinPath   string = "bin/trojango"
	HysteriaBinPath   string = "bin/hysteria"
	NaiveProxyBinPath string = "bin/naiveproxy"
	DownloadBaseUrl   string = "https://github.com/trojanpanel/install-script/releases/download/v1.2.0"

	TrojanPanelCoreVersion = "v2.0.0"
)
