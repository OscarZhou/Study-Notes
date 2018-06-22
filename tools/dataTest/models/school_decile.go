package models

import (
	"github.com/jinzhu/gorm"
)

type SchoolDecile struct {
	gorm.Model
	SchoolID uint
	Year     int64
	Decile   int
}
