package services

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Config struct {
	Env          string             `required:"true"`
	ServicePort  string             `required:"true"`
	DatabaseURL  string             `required:"true"`
	XMLRemoteURL string             `required:"true"`
	Logger       *zap.SugaredLogger `ignored:"true"`
	PGClient     *gorm.DB           `ignored:"true"`
}
