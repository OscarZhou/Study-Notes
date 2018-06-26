package main

import (
	"Study-Notes/tools/dataTest/models"
	"fmt"

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
	var schoolAcademics []models.NCEASchoolAcademicQualification
	err = db.Find(&schoolAcademics).Error
	if err != nil {
		return err
	}
	count := 0
	for _, v := range schoolAcademics {
		var (
			// schoolAcademic   models.NZSchoolAcademic
			schoolDataByYear models.SchoolDataByYear
		)
		// query := `SELECT nz_school_academics.*
		// FROM nz_school_academics
		// JOIN info_nz_categories ON info_nz_categories.id=nz_school_academics.category_id
		// JOIN info_nz_performances ON info_nz_performances.id=nz_school_academics.performance_id
		// JOIN info_nz_year_levels ON info_nz_year_levels.id=nz_school_academics.year_level_id
		// JOIN info_nz_performance_values ON info_nz_performance_values.performance_id = info_nz_performances.id
		// WHERE nz_school_academics.school_id=` + strconv.FormatUint(uint64(v.SchoolID), 10) + `
		// 	AND info_nz_categories.name='National'
		// 	AND nz_school_academics.year=2017
		// 	AND info_nz_performances.name='Qualification'
		// 	AND info_nz_year_levels.level=13
		// 	AND info_nz_performance_values.name = 'University Entrance'`

		// if db.Raw(query).Scan(&schoolAcademic).RecordNotFound() {
		// 	continue
		// }

		if (v.YearLevel == 13 && v.Qualification == "University Entrance") ||
			(v.YearLevel == 12 && v.Qualification == "NCEA Level 2") ||
			(v.YearLevel == 11 && v.Qualification == "NCEA Level 1") {
			schoolDataByYear.SchoolID = v.SchoolID
			schoolDataByYear.Year = v.Year
			schoolDataByYear.YearLevel = v.YearLevel
			schoolDataByYear.PassRate = v.CumulativeAchievementRate

			err = db.Create(&schoolDataByYear).Error
			if err != nil {
				return err

			}
			count++
		}
	}

	fmt.Printf("total: %d\n", count)
	return nil
}
