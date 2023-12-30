package services

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func (config *Config) ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	databaseClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	config.PGClient = databaseClient
	if err != nil {
		log.Fatal("Error connecting database")
	}

}
