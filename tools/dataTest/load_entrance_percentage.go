package main

import (
	"Study-Notes/tools/dataTest/models"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func LoadEntrancePercentageData() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	db.DropTable(
		models.SchoolDataByYear{},
	)

	db.AutoMigrate(
		models.SchoolDataByYear{},
	)

	// db.LogMode(true)

	// query the entrance percentage based on school decile table data
	var schoolDeciles []models.SchoolDecile
	err = db.Find(&schoolDeciles).Error
	if err != nil {
		return err
	}

	for _, v := range schoolDeciles {
		var (
			schoolAcademic   models.NZSchoolAcademic
			schoolDataByYear models.SchoolDataByYear
		)
		query := `SELECT nz_school_academics.* 
		FROM nz_school_academics  
		JOIN info_nz_categories ON info_nz_categories.id=nz_school_academics.category_id 
		JOIN info_nz_performances ON info_nz_performances.id=nz_school_academics.performance_id 
		JOIN info_nz_year_levels ON info_nz_year_levels.id=nz_school_academics.year_level_id 
		JOIN info_nz_performance_values ON info_nz_performance_values.performance_id = info_nz_performances.id 
		WHERE nz_school_academics.school_id=` + strconv.FormatUint(uint64(v.SchoolID), 10) + ` 
			AND info_nz_categories.name='National' 
			AND nz_school_academics.year=2017 
			AND info_nz_performances.name='Qualification' 
			AND info_nz_year_levels.level=13 
			AND info_nz_performance_values.name = 'University Entrance'`

		if db.Raw(query).Scan(&schoolAcademic).RecordNotFound() {
			continue
		}
		schoolDataByYear.SchoolID = v.SchoolID
		schoolDataByYear.Year = 2017
		schoolDataByYear.AdmissionRate = schoolAcademic.CumulativeAchievementRate

		err = db.Create(&schoolDataByYear).Error
		if err != nil {
			return err
		}
	}

	fmt.Printf("total: %d\n", len(schoolDeciles))
	return nil
}
