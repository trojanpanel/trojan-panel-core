package constant

const (
	// LogPath 日志文件夹路径
	LogPath string = "logs"

	// ConfigPath 全局配置文件夹路径
	ConfigPath string = "config"
	// ConfigFilePath 全局配置文件路径
	ConfigFilePath string = "config/config.ini"

	XrayPath         string = "bin/xray"
	XrayFilePath     string = "bin/xray/config.json"
	TrojanGoPath     string = "bin/trojango"
	TrojanGoFilePath string = "bin/trojango/config.json"
	HysteriaPath     string = "bin/hysteria"
	HysteriaFilePath string = "bin/hysteria/config.json"

	// SaltKey 加密 盐
	SaltKey string = "well_very_funny!"

	GrpcPortXray     = "60087"
	GrpcPortTrojanGo = "60088"
	GrpcPortHysteria = "60089"
)
