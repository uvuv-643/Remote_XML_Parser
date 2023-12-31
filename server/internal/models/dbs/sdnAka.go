package dbs

import (
	"gorm.io/gorm"
)

type SDNAka struct {
	gorm.Model

	UID      int64 `gorm:"primaryKey"`
	Type     string
	Category string
	LastName string

	Items []SDNItem `gorm:"many2many:items_akas;" `
}
