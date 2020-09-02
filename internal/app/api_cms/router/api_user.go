package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_cms/config"
	"grape/internal/app/api_cms/handlers/user"
	"grape/internal/pkg/middleware/router"
)

func setupRouterCMSUser(engine *gin.Engine) {
	path := config.Path + "/cms/v1/user"
	//path
	g := engine.Group(path)
	g.Use(router.AuthMiddleWareCheckToken())
	// user
	g.GET("/user/refresh", user.RefreshGin) //刷新令牌
	//g.POST("/user/login", user.LoginGin)            //登录
	g.GET("/user/permissions", user.PermissionsGin)        //查询自己拥有的权限
	g.PUT("/user", user.UpdateUserGin)                     //更新用户信息
	g.POST("/user/register", user.RegisterGin)             //注册
	g.PUT("/user/change_password", user.ChangePasswordGin) //修改密码
	g.GET("/user/information", user.InformationGin)        //查询自己信息
}
