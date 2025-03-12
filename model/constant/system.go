package constant

const (
	LogDir = "logs/"
	BinDir = "bin/"

	SystemLogPath = LogDir + "trojan-core.log"

	XrayConfigDir = BinDir + "xray/"
	XrayLogPath   = LogDir + "xray.log"
	XrayConfigExt = ".json"

	SingBoxConfigDir = BinDir + "sing-box/"
	SingBoxLogPath   = LogDir + "sing-box.log"

	HysteriaConfigDir = BinDir + "hysteria/"
	HysteriaLogPath   = LogDir + "hysteria.log"
	HysteriaConfigExt = ".yaml"

	NaiveProxyConfigDir = BinDir + "naiveproxy/"
	NaiveProxyLogPath   = LogDir + "naiveproxy.log"
	NaiveProxyConfigExt = ".json"

	SystemVersion = "v3.0.0"
)
