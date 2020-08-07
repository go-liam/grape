package main

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api-www/core"
	"log"
	"os"
	"os/signal"
)

//const (
//	port       = ":7203"
//	serverName = "api-www"
//	path       = "/" + serverName
//)

func main() {
	println("run api port ", core.Port)
	engine := gin.New()
	// 设置路由
	core.SetupRouterMetric(engine)
	core.SetupRouter(engine)

	engine.Run(core.Port)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
