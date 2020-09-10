package jwt2

import "github.com/go-liam/util/conv"

type LoginResp struct {
	Token   string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	ID      string `json:"id"`
}

func LoginToken(userID int64, clientID int, loginFlag int) *LoginResp {
	o := new(LoginResp)
	o.Token, o.Refresh = Create(userID, clientID, loginFlag)
	o.ID = conv.Int64ToString(userID)
	// 以后加入 单点逻辑
	return o
}
