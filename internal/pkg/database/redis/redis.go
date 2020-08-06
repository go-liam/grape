package redis

import (
	"fmt"
	redis2 "github.com/go-liam/cache/redis"
	"grape/internal/pkg/config/env"
)

var Server = new(redis2.SvRedis) //cache.InCache

func init() {
	Server.URL = fmt.Sprintf("%s@%s:%s", env.RedisConfig.RedisPwd, env.RedisConfig.RedisHost, env.RedisConfig.RedisPort) // "Liam123@localhost:6379"
	Server.RedisPrefix = env.RedisConfig.RedisPrefix
	Server.NewCache(Server.URL, Server.RedisPrefix)

}
