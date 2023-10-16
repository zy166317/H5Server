package db

import (
	"github.com/go-redis/redis"
	"time"
)

func ConnectToRedis(dbRedis *Redis) (*redis.Client, error) {
	redis := redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         dbRedis.Host + ":" + dbRedis.Port,
		Password:     "",
		DB:           0,
		PoolSize:     dbRedis.PoolSize,
		MinIdleConns: dbRedis.MinIdleConns,
		MaxConnAge:   time.Duration(dbRedis.MaxConnAge),
		IdleTimeout:  time.Duration(dbRedis.IdleTimeout),
	})
	err := redis.Ping().Err()
	if err != nil {
		return nil, err
	}
	return redis, nil
}
