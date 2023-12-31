package dbs

import (
	"time"
)

type SDNDateOfBirth struct {
	UID         int64 `gorm:"primaryKey"`
	DateOfBirth time.Time
	MainEntry   bool
}
