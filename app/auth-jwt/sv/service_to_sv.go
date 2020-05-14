package sv

import (
	"github.com/gin-gonic/gin"
	"log"
	"grape/pkg/services/jwt"
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
	pwd := c.Query("pwd")
	token, exTime, uid := jwt.ServerGin.Login(name, pwd,0)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    RespLogin{Token: token, ExpiresTime: exTime, Status: 0,UserID: uid},
	})
}

func CheckUser(c *gin.Context) {
	//log.Printf("head:%+v\n",c.Request.Header)
	//info
	token := c.GetHeader(TokenName)
	info, err := jwt.ServerGin.Info(token)
	if err !=nil{
		log.Printf("[ERROR] CheckUser:%+v\n ",err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    info,
	})
}

func Refresh(c *gin.Context) {
	token := c.GetHeader(TokenName)
	info, _ := jwt.ServerGin.Refresh(token)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    info,
	})
}

func Logout(c *gin.Context) {
	token := c.GetHeader(TokenName)
	jwt.ServerGin.Logout(token)
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
	if err != nil{
		log.Printf("[ERROR] CheckBadUser: %+v\n",err)
	}
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
