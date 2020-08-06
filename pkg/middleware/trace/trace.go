package trace

import (
	"github.com/go-liam/tracing/config"
)

//func init()  {
//	tracing.Init(&config.TraceConfig{IsOpen: true, HostPort: "127.0.0.1:6831", SamplerType: "const", SamplerParam: 0.01, LogSpans: true})
//}

var DefaultConfig *config.TraceConfig

func init() {
	DefaultConfig = &config.TraceConfig{IsOpen: true, HostPort: "127.0.0.1:6831", SamplerType: "const", SamplerParam: 0.01, LogSpans: true}
}
