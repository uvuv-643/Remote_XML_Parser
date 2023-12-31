package services

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Config struct {
	Env           string             `required:"true"`
	ServicePort   string             `required:"true"`
	DatabaseURL   string             `required:"true"`
	RedisURL      string             `required:"true"`
	RedisPassword string             `required:"true"`
	XMLRemoteURL  string             `required:"true"`
	Logger        *zap.SugaredLogger `ignored:"true"`
	PGClient      *gorm.DB           `ignored:"true"`
	RedisClient   *redis.Client      `ignored:"true"`
}
