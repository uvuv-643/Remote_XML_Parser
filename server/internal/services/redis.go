package services

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"reflect"
	"time"
)

func (config *Config) ConnectRedis() {
	config.RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisURL,
		Password: config.RedisPassword,
		DB:       0,
	})
}

func (config *Config) CacheGetOrSet(redisKey string, target interface{}) bool {
	redisResponse := config.RedisClient.Get(redisKey)
	if result, err := redisResponse.Result(); !errors.Is(err, redis.Nil) {
		redisType := reflect.TypeOf(target)
		redisNewElem := reflect.New(redisType)
		redisResponseObject := redisNewElem.Interface()
		err = json.Unmarshal([]byte(result), redisResponseObject)
		if err == nil {
			if reflect.DeepEqual(reflect.ValueOf(redisResponseObject).Elem().Interface(), target) {
				return true
			}
		}
	}
	jsonEncodedTarget, err := json.Marshal(target)
	if err != nil {
		return false
	}
	config.RedisClient.Set(redisKey, string(jsonEncodedTarget), time.Duration(int(time.Second)*config.RedisTTL))
	return false
}
