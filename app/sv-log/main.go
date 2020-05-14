package main

import (
	"log"
	"grape/app/sv-log/server"
	"os"
	"os/signal"
)

func main() {
	println(":run server")
	go server.RunServerGin()
	server.ServerGRPC(server.PortGRPC)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
