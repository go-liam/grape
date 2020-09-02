package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_cms/config"
	"grape/internal/app/api_cms/handlers/demo"
	"grape/internal/app/api_cms/handlers/user"
	"grape/internal/pkg/middleware/router"
	"net/http"
)

func setupRouterAPI(engine *gin.Engine) {
	path := config.Path + "/api/v1"
	engine.GET("/", Index)
	engine.GET(path, Index)
	//other
	g := engine.Group(path)
	g.POST("/login", user.LoginGin) //登录
	g.GET("/demos", demo.GetListGin)
}

func setupRouterCMS(engine *gin.Engine) {
	path := config.Path + "/cms/v1"
	engine.GET(path, Index)
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

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-Hello,It works.index "+config.ServerName)
}
