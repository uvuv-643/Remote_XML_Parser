package services

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

func (config *Config) ConnectRedis() {
	config.RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisURL,
		Password: config.RedisPassword,
		DB:       0,
	})
}

func (config *Config) CacheHit(redisKey string, target interface{}) bool {
	redisResponse := config.RedisClient.Get(redisKey)
	if redisResponse != nil {
		var redisResponseObject interface{}
		err := json.Unmarshal([]byte(redisResponse.String()), &redisResponseObject)
		if err == nil {
			if redisResponseObject == target {
				return true
			}
		}
	}
	return false
}
