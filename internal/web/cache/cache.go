package cache

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/ljinf/meet_server/pkg/config"
)

type Cache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) *Cache {
	return &Cache{
		rdb: rdb,
	}
}

func NewRedis(conf *config.AppConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.RedisDB.Host, conf.RedisDB.Port),
		Password: conf.RedisDB.Password,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}
