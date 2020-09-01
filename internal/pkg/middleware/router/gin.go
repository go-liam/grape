package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/pkg/util/jwt2"
	"net/http"
	"strings"
)

func AccessControlAllow() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Headers", "access-control-allow-headers,access-control-allow-methods,access-control-allow-origin,cache-control,content-type,utoken,tokenjwt,Authorization")
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,PATCH")
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusNoContent)
			c.Abort()
			return
		}
		c.Next()
	}
}

func AccessOPTIONSlAllow() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusNoContent)
			c.Abort()
			return
		}
		c.Next()
	}
}

// 过滤空头
func AuthMiddleWareCheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		errCode := 401
		if token == "" {
			c.Status(errCode)
			c.Abort()
			return
		}
		c.Next()
	}
}

// 取出用户信息
func GetJWTInfoByHeader(c *gin.Context) (*jwt2.CustomClaims, int) {
	token := c.Request.Header.Get("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	return jwt2.Server.ParseToken(token)
}
