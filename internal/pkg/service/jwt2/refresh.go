package jwt2

type RefreshResp struct {
	Token   string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func RefreshToken(refreshToken string) *RefreshResp {
	o := new(RefreshResp)
	o.Token, o.Refresh = Refresh(refreshToken)
	// 将来可以加上 单点登录逻辑
	return o
}
