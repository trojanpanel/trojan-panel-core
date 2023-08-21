package constant

const (
	CodeSuccess           int = 20000 // success
	CodeSysError          int = 50000 // system error
	CodeTokenExpiredError int = 50014 // login expired
	CodeIllegalTokenError int = 50008 // authentication failed
	CodeUnauthorizedError int = 50401 // not authentication
	CodeForbiddenError    int = 50403 // permission denied
)
