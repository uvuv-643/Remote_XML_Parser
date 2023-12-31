package dbs

import (
	"gorm.io/gorm"
)

type SDNAddress struct {
	gorm.Model

	UID             int64 `gorm:"primaryKey"`
	City            string
	Address1        string
	Address2        string
	Address3        string
	StateOrProvince string
	PostalCode      string
	Country         string

	Items []SDNItem `gorm:"many2many:items_addresses;"`
}
