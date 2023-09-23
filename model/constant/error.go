package constant

const (
	SysError            string = "system error, please contact the system administrator"
	ValidateFailed      string = "invalid argument"
	UsernameOrPassError string = "wrong username or password"

	UnauthorizedError string = "not authentication"
	IllegalTokenError string = "authentication failed"
	ForbiddenError    string = "permission denied"
	TokenExpiredError string = "login expired"

	GrpcError            string = "gRPC connection error"
	HttpError            string = "http connection error"
	XrayStartError       string = "failed to start xray"
	TrojanGoStartError   string = "failed to start trojango"
	HysteriaStartError   string = "failed to start hysteria"
	Hysteria2StartError  string = "failed to start hysteria2"
	NaiveProxyStartError string = "failed to start naiveproxy"
	ProcessStopError     string = "process suspend failed"

	NodeTypeNotExist   string = "node type does not exist"
	RemoveFileError    string = "failed to delete file"
	BinaryFileNotExist string = "binary file does not exist"
	ConfigFileNotExist string = "config file does not exist"

	GetLocalIPError string = "failed to obtain local IP"

	PortIsOccupied string = "the port is occupied, please check the port or choose another port"
	PortRangeError string = "the port range is between 100-30000"
)
