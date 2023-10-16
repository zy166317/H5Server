package db

import (
	"github.com/go-redis/redis"
)

var DBRedis *redis.Client

func RedisInit(redis *Redis) error {
	client, err := ConnectToRedis(redis)
	if err != nil {
		return err
	}
	DBRedis = client
	return nil
}
