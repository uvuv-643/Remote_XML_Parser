package dbs

import (
	"gorm.io/gorm"
)

type SDNPlaceOfBirth struct {
	gorm.Model

	UID       int64 `gorm:"primaryKey"`
	Place     string
	MainEntry bool

	Items []SDNItem `gorm:"many2many:items_places;"`
}
