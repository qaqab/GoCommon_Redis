package GoCommon_Redis

import (
	"github.com/go-redis/redis"
)

type RedisSetting struct {
	Addresse string `json:"addresse"`
	Password string `json:"password"`
	DB       int    `json:"db"`

	RedisClient *redis.Client `json:"client"`
}
