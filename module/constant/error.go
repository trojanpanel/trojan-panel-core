package constant

const (
	SysError string = "系统错误,请联系统管理员"

	GrpcError string = "gRPC连接异常"

	BinaryFileNotExist string = "二进制文件不存在"
	ConfigFileNotExist string = "配置文件不存在"

	NewXrayProcessError     string = "新建Xray进程失败"
	XrayStartError          string = "启动Xray失败"
	NewTrojanGoProcessError string = "新建TrojanGo进程失败"
	TrojanGoStartError      string = "启动TrojanGo失败"
	NewHysteriaProcessError string = "新建Hysteria进程失败"
	HysteriaStartError      string = "启动Hysteria失败"

	ProcessStopError string = "进程暂停失败"

	DownloadFilError string = "远程文件下载失败"

	RemoveFileError string = "删除文件失败"
)
