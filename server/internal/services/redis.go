package services

import (
	"github.com/go-redis/redis"
)

func (config *Config) ConnectRedis() {
	config.RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisURL,
		Password: config.RedisPassword,
		DB:       0,
	})
}
