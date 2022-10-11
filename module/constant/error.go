package constant

const (
	SysError            string = "系统错误,请联系统管理员"
	ValidateFailed      string = "参数无效"
	UsernameOrPassError string = "用户名或密码错误"

	UnauthorizedError string = "未认证"
	IllegalTokenError string = "认证失败"
	ForbiddenError    string = "没有权限"
	TokenExpiredError string = "登录过期"

	GrpcError          string = "gRPC连接异常"
	XrayStartError     string = "启动Xray失败"
	TrojanGoStartError string = "启动TrojanGo失败"
	HysteriaStartError string = "启动Hysteria失败"
	ProcessStopError   string = "进程暂停失败"

	NodeTypeNotExist   string = "节点类型不存在"
	RemoveFileError    string = "删除文件失败"
	BinaryFileNotExist string = "二进制文件不存在"
	ConfigFileNotExist string = "配置文件不存在"

	GetLocalIPError string = "获取本地IP失败"

	PortIsOccupied string = "端口被占用,请检查该端口或选择其他端口"
	PortRangeError string = "端口范围在100-30000之间"
)
