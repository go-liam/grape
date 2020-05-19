package cookie

import (
	"errors"
	"github.com/gin-gonic/gin"
	"grape/pkg/util/conv"
	"grape/pkg/util/jwt"
	"time"
)

/*
 maxAge int, 单位秒，有效期
			// path,cookie所在目录
			// domain string,域名
			// secure 是否能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
*/
type SvGin struct {
	TokenName    string
	TokenTime    int32
	CookieDomain string
	BadUidList   []int64
}

var ServerGin *SvGin

func init() {
	ServerGin = new(SvGin)
	ServerGin.TokenName = "tokenjwt"
	ServerGin.TokenTime = 3600 * 24 * 365
	ServerGin.CookieDomain = "mygs.com"
	ServerGin.BadUidList = []int64{1, 2, 3, 119} // 黑名单
}

func (sv *SvGin) Login(c *gin.Context, name string, userID int64) (int, error) {
	claims := &jwt.CustomClaims{
		UserID: userID,
		Name:   name,
	}
	minute := 10
	claims.NotBefore = time.Now().Unix() - 10000 // 签名生效时间
	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(minute)).Unix()
	claims.Issuer = jwt.TokenIssuer //签名的发行者
	token, _ := jwt.Server.CreateToken(claims)
	c.SetCookie(sv.TokenName, token, int(sv.TokenTime), "/", sv.CookieDomain, false, false)
	return 1, nil
}

func (sv *SvGin) Info(c *gin.Context) (*jwt.CustomClaims, error){
	cookie, err := c.Cookie(sv.TokenName)
	if err != nil {
		return nil,err
	}
	return jwt.Server.ParseToken(cookie)
}

//Refresh
func (sv *SvGin) Refresh(c *gin.Context ) (*jwt.CustomClaims, error) {
	cookie, err := c.Cookie(sv.TokenName )
	if err != nil {
		return nil,err
	}
	//info, _ := jwt.Server.ParseToken(cookie)
	//if info.ExpiresAt - time.Now().Unix() < 3600*24*7 {
	refresh, _ := jwt.Server.RefreshToken(cookie, sv.TokenTime)
	c.SetCookie(sv.TokenName, refresh, int(sv.TokenTime), "/", sv.CookieDomain, false, false)
	//}
	return jwt.Server.ParseToken(refresh)
}

func (sv *SvGin) Logout(c *gin.Context )  {
	c.SetCookie(sv.TokenName, "", -1000, "/", sv.CookieDomain, false, false)
}

// IsBadUser
func (sv *SvGin) CheckBadUser(c *gin.Context ) (int, *jwt.CustomClaims, error)  {
	cookie, err := c.Cookie(sv.TokenName )
	if err != nil {
		return 0,nil,err
	}
	//info
	info, _ := jwt.Server.ParseToken(cookie)
	flag := conv.ArrayInt64Contains(sv.BadUidList,info.UserID)
	if flag {
		return 4,nil , errors.New("bad uid")
	}
	return 1,info, nil
}