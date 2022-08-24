package middleware

//// JWTHandler jwt认证中间件
//func JWTHandler() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// Token放在Header的Authorization中，并使用Bearer开头
//		authHeader := c.Request.Header.Get("Authorization")
//		if authHeader == "" {
//			vo.Fail(constant.UnauthorizedError, c)
//			c.Abort()
//			return
//		}
//		// 按空格分割
//		parts := strings.SplitN(authHeader, " ", 2)
//		if !(len(parts) == 2 && parts[0] == "Bearer") {
//			vo.Fail(constant.IllegalTokenError, c)
//			c.Abort()
//			return
//		}
//		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
//		myClaims, err := util.ParseToken(parts[1])
//		if err != nil {
//			vo.Fail(err.Error(), c)
//			c.Abort()
//			return
//		}
//		get := redis.Client.String.
//			Get(fmt.Sprintf("trojan-panel:token:%s", myClaims.UserVo.Username))
//		result, err := get.String()
//		if err != nil || result == "" {
//			vo.Fail(constant.IllegalTokenError, c)
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
