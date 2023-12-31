package dbs

import (
	"time"
)

type SDNId struct {
	UID       int64 `gorm:"primaryKey"`
	Type      string
	Number    string
	Country   string
	IssueDate time.Time
}
