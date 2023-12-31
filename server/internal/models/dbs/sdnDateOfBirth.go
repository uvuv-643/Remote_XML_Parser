package dbs

import (
	"gorm.io/gorm"
	"time"
)

type SDNDateOfBirth struct {
	gorm.Model

	UID         int64 `gorm:"primaryKey"`
	DateOfBirth time.Time
	MainEntry   bool

	Items []SDNItem `gorm:"many2many:items_dates;"`
}
