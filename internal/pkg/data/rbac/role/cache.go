package role

import (
	"encoding/json"
	"fmt"
	"grape/configs/env"
	"grape/internal/pkg/database/redis"
)

const redisSecondTime = 32

func (e *SrvRole) CacheMulti() ([]*Model, error) {
	key := fmt.Sprintf("%s_rb_role_ls", env.RedisConfig.RedisPrefix)
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
	got, _ = e.FindMultiByNil()
	if len(got) > 0 {
		byteValue, _ := json.Marshal(&got)
		redis.Server.Set(key, byteValue, redisSecondTime)
	}
	return got, nil
}
