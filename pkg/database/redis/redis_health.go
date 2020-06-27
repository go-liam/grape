package redis

import (
	"log"
	"time"
)

func Health() string {
	log.Println("[INFO] redis url=",Server.URL)
	log.Println("[INFO] redis RedisPrefix=",Server.RedisPrefix)
	n := Server.RedisPrefix+"|health|0"
	Server.Set(n, time.Now().Unix(), 2)
	f,err := Server.Get(n)
	if err != nil{
		log.Printf("[ERROR] redis health:%+v\n",err)
		Server.NewCache(Server.URL, Server.RedisPrefix) //重新连接
	}
	return f
}
