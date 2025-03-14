package constant

const (
	LogDir = "logs/"
	BinDir = "bin/"

	SystemLogPath = LogDir + "trojan-core.log"

	XrayConfigDir = BinDir + "xray/"
	XrayLogPath   = LogDir + "xray.log"
	XrayConfigExt = ".json"

	SingBoxConfigDir = BinDir + "singbox/"
	SingBoxLogPath   = LogDir + "singbox.log"
	SingBoxConfigExt = ".json"

	HysteriaConfigDir = BinDir + "hysteria/"
	HysteriaLogPath   = LogDir + "hysteria.log"
	HysteriaConfigExt = ".yaml"

	NaiveProxyConfigDir = BinDir + "naiveproxy/"
	NaiveProxyLogPath   = LogDir + "naiveproxy.log"
	NaiveProxyConfigExt = ".json"

	SystemVersion = "v3.0.0"
)
