package data

import (
	"github.com/cyl615123/geek/week4/internal/config"
	"github.com/gomodule/redigo/redis"
)

type redisClient struct {
	*redis.Pool
}

func newRedis(cfg *config.RedisConf) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.Addr)
		},
	}
}

func (rc *redisClient) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := rc.Get()
	defer conn.Close()
	return rc.Do(commandName, args...)
}
