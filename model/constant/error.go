package constant

const (
	SysError          string = "system error"
	InvalidError      string = "invalid"
	UnauthorizedError string = "unauthorized"
	ForbiddenError    string = "permission denied"

	IllegalTokenError string = "authentication failed"

	GrpcError string = "failed to connect to remote service"
	HttpError string = "http connection error"
)
