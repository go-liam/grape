package core

import (
	"github.com/gin-gonic/gin"
	"grape/pkg/service/rbac/power"
	"grape/pkg/service/rbac/user"
	"grape/pkg/util/conv"
	"net/http"
)
var serverName = "auth-rbac"

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    serverName,
	})
}

func UserInfo(c *gin.Context) {
	userID := conv.StringToInt64(c.Param("userID"), 0)
	info, _ := user.Info(userID)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    info,
	})
}

func show(c *gin.Context,msg string, data interface{} ) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg ,
		"data":    data,
	})
}

func CheckUserInfo(c *gin.Context) {
	powerSt := c.Query("power")
	list ,_ := power.Server.List(1)
	pid := power.GetPowerIDByTag(powerSt, list)
	userID := conv.StringToInt64(c.Param("userID"), 0)
	info, _ := user.Info(userID)
	if info.RoleFlag == user.ConstRoleRoot {
		show(c,"超级管理员",info)
		return
	}
	flag := conv.ArrayIntContains(info.PowerIDs,pid)
	if flag{
		show(c,"有权限",info)
	}else {
		show(c,"没有权限",info)
	}
}
