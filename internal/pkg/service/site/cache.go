package site

import (
	"encoding/json"
	"fmt"
	"grape/configs/env"
	"grape/internal/pkg/database/redis"
)

const redisSecondTime = 32

func (e *SrvSite) CacheOne(id int64) (*Model, error) {
	key := fmt.Sprintf("%s_site_%d", env.RedisConfig.RedisPrefix, id)
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

func (e *SrvMenu) CacheMulti() ([]*MenuModel, error) {
	key := fmt.Sprintf("%s_menu_ls", env.RedisConfig.RedisPrefix)
	v, err := redis.Server.GetBytes(key)
	var got []*MenuModel
	var err2 error
	if err == nil {
		err2 = json.Unmarshal(v, &got)
		if err2 == nil {
			//println("redis data")
			return got, nil
		}
	}
	got, _ = e.FindMulti()
	if got != nil && len(got) > 0 {
		byteValue, _ := json.Marshal(&got)
		redis.Server.Set(key, byteValue, redisSecondTime)
	}
	return got, nil
}
