package redis

import "github.com/alicebob/miniredis/v2"

func MockRedis() {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	Server.URL = s.Addr()
	Server.NewCache(Server.URL, Server.RedisPrefix)
	//RedisClient = redis.NewClient(&redis.Options{
	//	Addr: s.Addr(),
	//})
	//_, err = RedisClient.Ping().Result()
	//if err != nil {
	//	panic(err)
	//}
}
