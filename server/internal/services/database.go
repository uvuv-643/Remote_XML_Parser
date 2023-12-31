package services

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func (config *Config) ConnectDatabase() {
	databaseClient, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})
	config.PGClient = databaseClient
	if err != nil {
		log.Fatal("Error connecting database")
	}
}
