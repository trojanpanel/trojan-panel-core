package constant

// 状态码
const (
	CodeSuccess           int = 20000 // 成功
	CodeSysError          int = 50000 // 系统错误,请联系统管理员
	CodeTokenExpiredError int = 50014 // 登录过期
	CodeIllegalTokenError int = 50008 // 认证失败
	CodeUnauthorizedError int = 50401 // 未认证
	CodeForbiddenError    int = 50403 // 暂无权限
)
