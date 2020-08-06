package single

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"grape/pkg/service/jwt"
	jwt2 "grape/pkg/util/jwt"
	"log"
	"net/http"
)

const TokenName = "tokenjwt"

type RespLogin struct {
	Status      int32  `json:"status"`
	Token       string `json:"token"`
	ExpiresTime int32  `json:"expiresTime"`
	UserID      int64  `json:"userID"`
}

func Login(c *gin.Context) {
	name := c.Query("name")
	userID := c.Query("userID")
	clientID := conv.StringToInt(c.Query("clientID"), 0)
	token, exTime, uid := jwt.ServerGin.Login(name, userID, clientID)
	if uid > 0 {
		jwt.ServerMock.UpdateOrAdd(&jwt.ModelUserToken{ClientID: clientID, UserID: uid, Token: token})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    RespLogin{Token: token, ExpiresTime: exTime, Status: 0, UserID: uid},
	})
}

func CheckUser(c *gin.Context) {
	log.Printf("head:%+v\n", c.Request.Header)
	//info
	token := c.GetHeader(TokenName)
	info, _ := jwt.ServerGin.Info(token)
	if info == nil || conv.StringToInt64(info.UserID, 0) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "JWT有误",
			"data":    nil,
		})
		return
	}
	info2, _ := jwt.ServerMock.GET(conv.StringToInt64(info.UserID, 0), info.ClientType)
	if info2 == nil || info2.Token != token {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Token 已经被退出",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    info2,
	})
}

func Refresh(c *gin.Context) {
	token := c.GetHeader(TokenName)
	info, _ := jwt.ServerGin.Info(token)
	if info == nil || info.UserID == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "JWT数据有误",
			"data":    nil,
		})
		return
	}
	info2, _ := jwt.ServerMock.GET(conv.StringToInt64(info.UserID, 0), info.ClientType)
	if info2 == nil || info2.UserID <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Redis没有此用户数据",
			"data":    nil,
		})
		return
	}
	refresh, _ := jwt.ServerGin.Refresh(token)
	jwt.ServerMock.UpdateOrAdd(&jwt.ModelUserToken{ClientID: info.ClientType, UserID: conv.StringToInt64(info.UserID, 0), Token: refresh})
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    refresh,
	})
}

func Logout(c *gin.Context) {
	token := c.GetHeader(TokenName)
	jwt.ServerGin.Logout(token)
	info, _ := jwt2.Server.ParseToken(token)
	log.Printf("jwt:%+v\n", info)
	if info == nil || info.UserID == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "JWT无效",
			"data":    "logout",
		})
		return
	}
	jwt.ServerMock.DeleteByUserIDAndClient(conv.StringToInt64(info.UserID, 0), info.ClientType)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    "logout",
	})
}

func CheckBadUser(c *gin.Context) {
	//info
	token := c.GetHeader(TokenName)
	f, info, err := jwt.ServerGin.IsBadUser(token)
	if f {
		c.JSON(http.StatusOK, gin.H{
			"code":    119,
			"message": "黑名单",
			"data":    err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    info,
	})
}
