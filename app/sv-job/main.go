package main

import (
	"grape/app/sv-job/core"
	"log"
	"os"
	"os/signal"
)

func main() {
	println(":run server")
	//go server.RunServerGin()
	core.JobHour()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
