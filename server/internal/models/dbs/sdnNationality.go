package dbs

import (
	"gorm.io/gorm"
)

type SDNNationality struct {
	gorm.Model

	UID       int64 `gorm:"primaryKey"`
	Country   string
	MainEntry bool

	Items []SDNItem `gorm:"many2many:items_nationalities;"`
}
