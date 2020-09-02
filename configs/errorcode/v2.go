package errorcode

const (
	// 1004* 给 token ; 1005* 给 refreshToken
	RefreshTokenExpired    = 10052
	RefreshTokenExpiredMsg = "RefreshToken过期"
	// login : 1001* 给登录
	LoginError    = 10010
	LoginErrorMsg = "登录不成功"
)
