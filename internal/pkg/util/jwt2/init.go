package jwt2

import (
	"fmt"
	"log"
	"time"
)

var Server = NewJWT(SignKey)
var ServerRefresh = NewJWT(SignKeyRefresh)

const (
	// SignKey ：加密盐,每段时间更新盐:新旧盐可同时进行
	SignKey        = "188274e3052a9fd6450077a3287003a-mb1907@jm@!$$v2" //Token
	SignKeyRefresh = "388274e3052a9fd6450077a3287003a-mb1907@jm@!$$v3" //Refresh
	RefreshOutTime = 30 * 24 * 3600                                    // 30天
	TokenOutTime   = 5 * 60                                            // 5分钟
)

func Create(userID int64, clientType int) (string, string) {
	claims1 := &CustomClaims{
		UserID:     fmt.Sprintf("%d", userID),
		ClientType: clientType,
	}
	claims1.ExpiresAt = time.Now().Add(time.Second * time.Duration(RefreshOutTime)).Unix()
	refresh, _ := ServerRefresh.CreateToken(claims1)
	claims2 := &CustomClaims{
		UserID:     fmt.Sprintf("%d", userID),
		ClientType: clientType,
	}
	claims2.ExpiresAt = time.Now().Add(time.Second * time.Duration(TokenOutTime)).Unix()
	token, _ := Server.CreateToken(claims2)
	return token, refresh
}

func Refresh(tokenString string, refreshString string) (string, string) {
	claims1, err2 := ServerRefresh.ParseToken(refreshString)
	if err2 != Success {
		log.Printf("[ERROR] Refresh F: %+v\n", err2)
		return "", ""
	}
	claims1.ExpiresAt = time.Now().Add(time.Second * time.Duration(RefreshOutTime)).Unix()
	refresh, _ := ServerRefresh.CreateToken(claims1)

	claimsT, errT := Server.ParseToken(tokenString)
	if errT != Success {
		log.Printf("[ERROR] Refresh T: %+v\n", errT)
		return "", ""
	}
	claimsT.ExpiresAt = time.Now().Add(time.Second * time.Duration(TokenOutTime)).Unix()
	token, _ := Server.CreateToken(claimsT)
	return token, refresh
}
