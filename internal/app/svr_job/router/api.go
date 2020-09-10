package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/svr_job/config"
	"grape/internal/app/svr_job/handlers/health"
)

func SetupRouter(engine *gin.Engine) {
	path := config.Path + "/api/v1"
	engine.GET("/", health.Index)
	engine.GET(path, health.Index)
}
