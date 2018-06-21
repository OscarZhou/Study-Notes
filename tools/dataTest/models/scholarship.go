package models

import (
	"github.com/jinzhu/gorm"
)

type Scholarship struct {
	gorm.Model
	Year                    int64
	School                  string
	Scholarships            int64
	OutstandingScholarships int64
	SchoolID                uint
}
