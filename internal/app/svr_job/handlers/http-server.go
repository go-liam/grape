package handlers

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/svr_job/config"
	"grape/internal/app/svr_job/router"
)

func RunHttpServer() {
	engine := gin.New()
	// 设置路由
	router.SetupRouter(engine)
	engine.Run(config.Port)
}
