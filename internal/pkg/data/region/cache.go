package region

import (
	"encoding/json"
	"fmt"
	"grape/configs/env"
	"grape/internal/pkg/database/redis"
)

const redisSecondTime = 32

func (e *SrvRegion) CacheOne(uid int64) (*Model, error) {
	key := fmt.Sprintf("%s_region_%d", env.RedisConfig.RedisPrefix, uid)
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
	info, _ = e.FindOne(uid)
	if info != nil && (info.UserID) > 0 {
		byteValue, _ := json.Marshal(&info)
		redis.Server.Set(key, byteValue, redisSecondTime)
	}
	return info, nil
}
