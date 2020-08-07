package main

import (
	"github.com/go-liam/tracing/config"
	"github.com/go-liam/tracing/jaeger"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"grape/configs"
	"grape/internal/app/gateway-api/core"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	println(configs.ServerGateWayAPI.Name, " running on port:", configs.ServerGateWayAPI.Port)
	//链路跟踪
	sv := new(jaeger.SvJeager)
	sv.Init(config.DefaultConfig)
	tracer, closer, _ := sv.NewJaegerTracer("gateway-api", sv.Config().HostPort)
	defer closer.Close()
	core.DateInit()
	// start server
	http.HandleFunc("/", core.HandleRequestAndRedirect)
	if err := http.ListenAndServe(configs.ServerGateWayAPI.Port, nethttp.Middleware(tracer, http.DefaultServeMux)); err != nil {
		log.Printf("[ERROR] %+v\n", err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
