package services

import (
	"Remote_XML_Parser/internal/models/dbs"
	"Remote_XML_Parser/internal/models/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func (config *Config) ConnectDatabase() {
	databaseClient, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	config.PGClient = databaseClient
	if err != nil {
		log.Fatal("Error connecting database")
	}
}

func (config *Config) MigrateDatabase() {

	config.PGClient.AutoMigrate(
		&dbs.SDNAddress{},
		&dbs.SDNAka{},
		&dbs.SDNCitizenship{},
		&dbs.SDNDateOfBirth{},
		&dbs.SDNId{},
		&dbs.SDNNationality{},
		&dbs.SDNPlaceOfBirth{},
		&dbs.SDNProgram{},
		&dbs.SDNItem{},
		&global.UpdateStatus{},
	)
}
