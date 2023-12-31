package dbs

import (
	"gorm.io/gorm"
	"time"
)

type SDNId struct {
	gorm.Model

	UID       int64 `gorm:"primaryKey"`
	Type      string
	Number    string
	Country   string
	IssueDate time.Time

	Items []SDNItem `gorm:"many2many:items_aids;"`
}
