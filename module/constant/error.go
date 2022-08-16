package constant

const (
	SysError            string = "系统错误,请联系统管理员"
	ValidateFailed      string = "参数无效"
	UsernameOrPassError string = "用户名或密码错误"

	GrpcError               string = "gRPC连接异常"
	NewXrayProcessError     string = "新建Xray进程对象失败"
	XrayStartError          string = "启动Xray协程失败"
	NewTrojanGoProcessError string = "新建TrojanGo进程对象失败"
	TrojanGoStartError      string = "启动TrojanGo协程失败"
	NewHysteriaProcessError string = "新建Hysteria进程对象失败"
	HysteriaStartError      string = "启动Hysteria协程失败"
	ProcessStopError        string = "进程暂停失败"

	DownloadFilError   string = "远程文件下载失败"
	RemoveFileError    string = "删除文件失败"
	BinaryFileNotExist string = "二进制文件不存在"
	ConfigFileNotExist string = "配置文件不存在"
)
