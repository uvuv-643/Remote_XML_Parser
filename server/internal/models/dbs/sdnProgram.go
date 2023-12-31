package dbs

import "gorm.io/gorm"

type SDNProgram struct {
	gorm.Model
	Program string    `gorm:"primaryKey"`
	Items   []SDNItem `gorm:"many2many:items_programs;"`
}
