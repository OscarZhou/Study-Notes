package models

type SchoolDataByYear struct {
	ID        uint `gorm:"primary_key"`
	SchoolID  uint
	Year      int64
	YearLevel int64
	PassRate  float64
}
