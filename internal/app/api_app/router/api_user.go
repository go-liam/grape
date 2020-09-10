package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_app/config"
	"grape/internal/app/api_app/handlers/jwt"
	"grape/internal/app/api_app/handlers/user"
	"grape/internal/pkg/middleware/router"
)

func setupRouterUser(engine *gin.Engine) {
	path := config.Path + "/api/v1"
	g := engine.Group(path)
	// login
	g.POST("/user/login-pwd", user.LoginByPwdGin)
	g.GET("/user/login-sms", user.LoginMsmGin)  //发短信
	g.POST("/user/login-sms", user.LoginMsmGin) //登录
	g.GET("/user/logout", user.LogoutGin)
	// token
	g.Use(router.AuthMiddleWareCheckToken())
	// user
	g.GET("/user/refresh", jwt.RefreshGin) //刷新令牌
	g.GET("/user/info", user.InfoGin)
}
