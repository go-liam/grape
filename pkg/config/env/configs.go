package env

import (
	"github.com/timest/env"
"log"
)
const (
	ConstEnvLocal = "local"
	ConstEnvDev = "dev"
	ConstEnvTest = "test"
	ConstEnvProd = "prod"
	ConstEnvPre = "pre"
	ConstEnvUnit = "unit" // 单元测试环境
)

var (
	// EnvConfig :
	EnvConfig *config
)

type config struct {
	ProjectEnv string `env:"PROJECT_ENV" default:"dev"`
	APIVersion string `env:"API_VERSION" default:"Commit ID"`
}

func init() {
	EnvConfig = new(config)
	env.IgnorePrefix()
	err := env.Fill(EnvConfig)
	log.Printf("[INFO] config :%+v\n", EnvConfig)
	if err != nil {
		panic(err)
	}
}

