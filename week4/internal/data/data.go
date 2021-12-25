package data

import (
	"github.com/cyl615123/geek/week4/internal/config"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewBookRepo, NewData)

type Data struct {
	redis *redisClient
}

func NewData(cfg *config.Conf) (*Data, func(), error) {
	data := &Data{
		redis: newRedisClient(&cfg.Redis),
	}
	return data, func() {
		data.redis.Close()
	}, nil
}

func newRedisClient(cfg *config.RedisConf) *redisClient {
	return &redisClient{
		newRedis(cfg),
	}
}
