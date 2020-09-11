package server

import (
	"encoding/json"
	"fmt"
	"grape/configs/env"
	"grape/internal/pkg/database/redis"
)

const redisSecondTime = 32

func (e *SrvServer) CacheOne(id int64) (*Model, error) {
	key := fmt.Sprintf("%s_server_%d", env.RedisConfig.RedisPrefix, id)
	v, err := redis.Server.GetBytes(key)
	var info *Model
	var err2 error
	if err == nil {
		err2 = json.Unmarshal(v, &info)
		if err2 == nil {
			//println("redis data")
			return info, nil
		}
	}
	info, _ = e.FindOne(id)
	if info != nil && (info.ID) > 0 {
		byteValue, _ := json.Marshal(&info)
		redis.Server.Set(key, byteValue, redisSecondTime)
	}
	return info, nil
}

func (e *SrvServer) CacheMulti() ([]*Model, error) {
	key := fmt.Sprintf("%s_servers", env.RedisConfig.RedisPrefix)
	got, err := e.getCacheMulti(key)
	if got != nil {
		return got, err
	}
	return e.setCacheMulti(key)
}

func (e *SrvServer) getCacheMulti(key string) ([]*Model, error) {
	v, err := redis.Server.GetBytes(key)
	var got []*Model
	var err2 error
	if err == nil {
		err2 = json.Unmarshal(v, &got)
		if err2 == nil {
			//println("redis data")
			return got, nil
		}
	}
	return nil, err2
}

func (e *SrvServer) setCacheMulti(key string) ([]*Model, error) {
	got, _ := e.FindMultiByNil()
	if got != nil && len(got) > 0 {
		byteValue, _ := json.Marshal(&got)
		redis.Server.Set(key, byteValue, redisSecondTime)
	}
	return got, nil
}
