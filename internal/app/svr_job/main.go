package main

import (
	"grape/internal/app/svr_job/handlers"
	"grape/internal/app/svr_job/handlers/job"
	"log"
	"os"
	"os/signal"
)

func main() {
	println(":run server")
	go handlers.RunHttpServer()
	job.RunJob()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
