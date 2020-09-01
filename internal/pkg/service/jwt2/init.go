package jwt2

import (
	"grape/internal/pkg/util/jwt2"
	"log"
	"time"
)

/*
	将来可以加上 单点登录逻辑（放在刷新逻辑）
*/
var Server = new(jwt2.JWT)
var ServerRefresh = new(jwt2.JWT)

var (
	// SignKey ：加密盐,每段时间更新盐:新旧盐可同时进行
	SignKey        = "188274e3052a9fd6450077a3287003a-mb1907@jm@!$$v2" //Token
	SignKeyRefresh = "388274e3052a9fd6450077a3287003a-mb1907@jm@!$$v3" //Refresh
	RefreshOutTime = 30 * 24 * 3600                                    // 30天
	TokenOutTime   = 30 * 60                                           // X 分钟(单点登录时调为 1分钟)
)

func init() {
	// key 可以从 env 取值
	Server.SetKey(SignKey)
	ServerRefresh.SetKey(SignKeyRefresh)
}

func Create(userID int64, clientType int, loginFlag int) (string, string) {
	claims1 := &jwt2.CustomClaims{
		UserID:     userID, // fmt.Sprintf("%d", userID),
		ClientType: clientType,
		LoginFlag:  loginFlag,
	}
	claims1.ExpiresAt = time.Now().Add(time.Second * time.Duration(RefreshOutTime)).Unix()
	refresh, _ := ServerRefresh.CreateToken(claims1)
	//token
	claims1.ExpiresAt = time.Now().Add(time.Second * time.Duration(TokenOutTime)).Unix()
	token, _ := Server.CreateToken(claims1)
	return token, refresh
}

func Refresh(refreshString string) (string, string) {
	var diff int64 = 10 * 24 * 3600
	claims1, err2 := ServerRefresh.ParseToken(refreshString)
	if err2 != jwt2.Success {
		log.Printf("[ERROR] Refresh F: %+v\n", err2)
		return "", ""
	}
	if claims1.ExpiresAt-diff <= time.Now().Unix() {
		println("[INFO] new refreshString")
		claims1.ExpiresAt = time.Now().Add(time.Second * time.Duration(RefreshOutTime)).Unix()
		refreshString, _ = ServerRefresh.CreateToken(claims1)
	}
	//token
	claims1.ExpiresAt = time.Now().Add(time.Second * time.Duration(TokenOutTime)).Unix()
	token, _ := Server.CreateToken(claims1)

	return token, refreshString
}
