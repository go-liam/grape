package jwt

import (
	"errors"
	"github.com/go-liam/util/conv"
	"grape/pkg/util/jwt"
	"time"
)

/*
 */
type SvGin struct {
	TokenName  string
	TokenTime  int32
	BadUidList []int64
}

var ServerGin *SvGin

func init() {
	ServerGin = new(SvGin)
	ServerGin.TokenName = "tokenjwt"
	ServerGin.TokenTime = 3600 * 24 * 365        // 秒
	ServerGin.BadUidList = []int64{1, 2, 3, 119} // 黑名单
	//ServerGin.
}

// Login :[token,退出时间（秒），err]
func (sv *SvGin) Login(name string, psw string, clientID int) (string, int32, int64) {
	uid := psw // int64( utf8.RuneCountInString(psw) )
	claims := &jwt.CustomClaims{
		UserID:     uid,
		Name:       name,
		ClientType: clientID,
	}
	claims.NotBefore = time.Now().Unix() - 10000 // 签名生效时间
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(sv.TokenTime)).Unix()
	claims.Issuer = jwt.TokenIssuer //签名的发行者
	token, _ := jwt.Server.CreateToken(claims)
	return token, sv.TokenTime, conv.StringToInt64(uid, 0)
}

func (sv *SvGin) Info(token string) (*jwt.CustomClaims, error) {
	return jwt.Server.ParseToken(token)
}

// Refresh :
func (sv *SvGin) Refresh(token string) (string, error) {
	return jwt.Server.RefreshToken(token, sv.TokenTime)
	//return jwt.Server.ParseToken(refresh)
}

// Logout :
func (sv *SvGin) Logout(token string) {
	println("logout:", token)
}

// IsBadUser :
func (sv *SvGin) IsBadUser(token string) (bool, *jwt.CustomClaims, error) {
	//info
	info, err := jwt.Server.ParseToken(token)
	if err != nil {
		return false, nil, errors.New("bad token")
	}
	flag := conv.ArrayInt64Contains(sv.BadUidList, conv.StringToInt64(info.UserID, 0))
	if flag {
		return true, nil, errors.New("bad uid")
	}
	return false, info, nil
}

// WasQuit :用户是否 单点登录，被退出
func (sv *SvGin) WasQuit(token string) (bool, error) {
	//info
	return false, nil
}
