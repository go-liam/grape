package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api-cms/config"
	"grape/internal/app/api-cms/handlers/user"
	"net/http"
)

func setupRouterAPI(engine *gin.Engine) {
	path := config.Path + "/api/v1"
	engine.GET("/", Index)
	engine.GET(path, Index)
	//other
	//g := engine.Group(path)
	//g.POST("/login", user.LoginGin)
}

func setupRouterCMS(engine *gin.Engine) {
	path := config.Path + "/cms/v1"
	engine.GET(path, Index)
	g := engine.Group(path)
	// user
	g.POST("/user/login", user.LoginGin)            //登录
	g.GET("/user/permissions", user.PermissionsGin) //查询自己拥有的权限
	g.PUT("/user", user.UpdateUserGin)              //更新用户信息
	g.POST("/user/register", Index)                 //注册
	g.PUT("/user/change_password", Index)           //修改密码
	g.GET("/user/refresh", Index)                   //刷新令牌
	g.GET("/user/information", Index)               //查询自己信息
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-Hello,It works.index "+config.ServerName)
}
