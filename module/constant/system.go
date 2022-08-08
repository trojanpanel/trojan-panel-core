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

	XrayPath           string = "bin/xray"
	XrayConfigFilePath string = "bin/xray/config.json"
	TrojanGoPath       string = "bin/trojango"
	HysteriaPath       string = "bin/hysteria"

	GrpcPortXray string = "60087"

	DownloadBaseUrl string = "https://github.com/trojanpanel/install-script/releases/download/v1.2.0/"
)
