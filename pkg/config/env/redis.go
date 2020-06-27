package env

import (
	en "github.com/timest/env"
	"log"
)

var (
	// MysqlConfig :
	RedisConfig *redisConfig
)

type redisConfig struct {
	// redis
	RedisHost   string `env:"REDIS_HOST" default:"localhost"`
	RedisPort   string `env:"REDIS_PORT" default:"6379"`
	RedisPwd    string `env:"REDIS_PWD" default:"block123"`
	RedisPrefix string `env:"REDIS_PREFIX" default:"b|"`
}

func init() {
	RedisConfig = new(redisConfig )
	en.IgnorePrefix()
	err := en.Fill(RedisConfig)
	log.Printf("[INFO] RedisConfig :%+v\n", RedisConfig)
	if err != nil {
		log.Printf("[ERROR] RedisConfig :%+v\n", err)
	}
}

