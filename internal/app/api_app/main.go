package main

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_app/config"
	"grape/internal/app/api_app/router"
	"log"
	"os"
	"os/signal"
)

func main() {
	println("run api port ", config.Port)
	engine := gin.New()
	// 设置路由
	router.SetupRouter(engine)
	router.SetupRouterMetric(engine)
	engine.Run(config.Port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
