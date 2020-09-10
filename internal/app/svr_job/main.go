package main

import (
	"grape/internal/app/svr_job/core"
	"grape/internal/app/svr_job/handlers"
	"log"
	"os"
	"os/signal"
)

func main() {
	println(":run server")
	go handlers.RunHttpServer()
	core.JobHour()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
