package main

import (
	"github.com/go-liam/tracing/config"
	"github.com/go-liam/tracing/jaeger"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"grape/internal/app/gateway-api/core"
	config2 "grape/pkg/config"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	println(config2.ServerGateWayAPI.Name, " running on port:", config2.ServerGateWayAPI.Port)
	//链路跟踪
	sv := new(jaeger.SvJeager)
	sv.Init(config.DefaultConfig)
	tracer, closer, _ := sv.NewJaegerTracer("gateway-api", sv.Config().HostPort)
	defer closer.Close()
	core.DateInit()
	// start server
	http.HandleFunc("/", core.HandleRequestAndRedirect)
	if err := http.ListenAndServe(config2.ServerGateWayAPI.Port,nethttp.Middleware(tracer, http.DefaultServeMux)); err != nil {
		log.Printf("[ERROR] %+v\n", err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
