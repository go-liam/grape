package jwt2

type LoginResp struct {
	Token   string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func LoginToken(userID int64, clientID int, loginFlag int) *LoginResp {
	o := new(LoginResp)
	o.Token, o.Refresh = Create(userID, clientID, loginFlag)
	// 以后加入 单点逻辑
	return o
}
