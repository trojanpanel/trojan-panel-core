package constant

const (
	// LogPath 日志文件夹路径
	LogPath string = "logs"

	// ConfigPath 全局配置文件夹路径
	ConfigPath string = "config"
	// ConfigFilePath 全局配置文件路径
	ConfigFilePath string = "config/config.ini"

	XrayPath                      = "bin/xray"
	XrayConfigFilePath     string = "bin/xray/config.json"
	TrojanGoPath                  = "bin/trojango"
	TrojanGoConfigFilePath string = "config/config.json"
	HysteriaPath                  = "bin/hysteria"
	HysteriaConfigFilePath string = "config/config.json"

	// SaltKey 加密 盐
	SaltKey string = "well_very_funny!"

	GrpcPortXray     string = "60087"
	GrpcPortTrojanGo string = "60088"
	GrpcPortHysteria string = "60089"
)
