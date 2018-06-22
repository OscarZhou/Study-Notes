package main

import (
	"Study-Notes/tools/dataTest/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func FillScholarshipYear() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	// db.DropTable(
	// 	models.NZCategory{},
	// )

	// db.AutoMigrate(
	// 	models.NZCategory{},
	// )

	// db.LogMode(true)

	var scholarships []models.Scholarship
	err = db.Find(&scholarships).Error
	if err != nil {
		return err
	}

	years := []int64{2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017}
	var schools []School
	err = db.Raw(`select school_id from scholarships group by school_id order by school_id asc`).Scan(&schools).Error
	if err != nil {
		return err
	}
	tx := db.Begin()
	// The last value is 0
	for _, school := range schools[:len(schools)-1] {
		for _, year := range years {
			var scholarship models.Scholarship
			if tx.Where("school_id = ? and year = ?", school.SchoolID, year).Find(&scholarship).RecordNotFound() {

				err = tx.Limit(1).Where("school_id = ?", school.SchoolID).Find(&scholarship).Error
				if err != nil {
					tx.Rollback()
					return err
				}

				var newScholarship models.Scholarship
				newScholarship.SchoolID = scholarship.SchoolID
				newScholarship.School = scholarship.School
				newScholarship.Year = year
				newScholarship.Scholarships = 0
				newScholarship.OutstandingScholarships = 0
				err = tx.Create(&newScholarship).Error
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}
	tx.Commit()

	fmt.Println("finish.")
	return nil
}

type School struct {
	SchoolID uint `json:"school_id"`
}
