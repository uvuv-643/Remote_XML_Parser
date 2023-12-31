package dbs

import (
	"gorm.io/gorm"
)

type SDNItem struct {
	gorm.Model
	UID              int64 `gorm:"primaryKey"`
	FirstName        string
	LastName         string
	Title            string
	SDNType          string
	Remarks          string
	Program          []SDNProgram      `gorm:"many2many:items_programs;"`
	Aka              []SDNAka          `gorm:"many2many:items_akas;"`
	ID               []SDNId           `gorm:"many2many:items_aids;"`
	Address          []SDNAddress      `gorm:"many2many:items_addresses;"`
	Nationality      []SDNNationality  `gorm:"many2many:items_nationalities;"`
	DateOfBirthItem  []SDNDateOfBirth  `gorm:"many2many:items_dates;"`
	PlaceOfBirthItem []SDNPlaceOfBirth `gorm:"many2many:items_places;"`
	Citizenship      []SDNCitizenship  `gorm:"many2many:items_citizenships;"`
}
