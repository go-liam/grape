package main

import (
	"log"
	"grape/app/sv-region/server"
	"os"
	"os/signal"
)

func main() {
	println(":run server")
	go server.RunServerGin()
	server.RunServerGRPC()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
