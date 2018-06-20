package main

import (
	"Study-Notes/tools/dataTest/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ImportData imports all information data
func ImportData() error {

	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	db.DropTable(
		&models.NZPerformance{},
		&models.NZPerformanceValue{},
		&models.NZCategory{},
		&models.NZYearLevel{},
	)

	db.AutoMigrate(
		&models.NZPerformance{},
		&models.NZPerformanceValue{},
		&models.NZCategory{},
		&models.NZYearLevel{},
	)

	err = addEndorsement(db)
	if err != nil {
		panic(err)
	}

	err = addQualification(db)
	if err != nil {
		panic(err)
	}

	err = addLit(db)
	if err != nil {
		panic(err)
	}

	err = addYearLevel(db)
	if err != nil {
		panic(err)
	}

	err = addGender(db)
	if err != nil {
		panic(err)
	}

	err = addEthnicity(db)
	if err != nil {
		panic(err)
	}

	err = addDecileBand(db)
	if err != nil {
		panic(err)
	}

	err = addNation(db)
	if err != nil {
		panic(err)
	}
	return nil
}

func addEndorsement(db *gorm.DB) error {
	var err error
	var performanceValues []models.NZPerformanceValue
	performanceValues = append(performanceValues, models.NZPerformanceValue{Title: "Qualification", Name: "University Entrance"})
	performanceValues = append(performanceValues, models.NZPerformanceValue{Title: "Qualification", Name: "NCEA Level 2"})
	performanceValues = append(performanceValues, models.NZPerformanceValue{Title: "Qualification", Name: "NCEA Level 3"})
	performanceValues = append(performanceValues, models.NZPerformanceValue{Title: "Qualification", Name: "NCEA Level 1"})

	var performanceValues2 []models.NZPerformanceValue
	performanceValues2 = append(performanceValues2, models.NZPerformanceValue{Title: "Endorsement", Name: "Excellence"})
	performanceValues2 = append(performanceValues2, models.NZPerformanceValue{Title: "Endorsement", Name: "Merit"})
	performanceValues2 = append(performanceValues2, models.NZPerformanceValue{Title: "Endorsement", Name: "No Endorsement"})

	for _, qualification := range performanceValues {
		for _, endorsement := range performanceValues2 {
			var performances = models.NZPerformance{Name: "Endorsement"}

			performances.PerformanceValue = append(performances.PerformanceValue, qualification)
			performances.PerformanceValue = append(performances.PerformanceValue, endorsement)

			err = db.Create(&performances).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func addQualification(db *gorm.DB) error {
	var err error
	var performanceValues []models.NZPerformanceValue
	performanceValues = append(performanceValues, models.NZPerformanceValue{Title: "Qualification", Name: "University Entrance"})
	performanceValues = append(performanceValues, models.NZPerformanceValue{Title: "Qualification", Name: "NCEA Level 2"})
	performanceValues = append(performanceValues, models.NZPerformanceValue{Title: "Qualification", Name: "NCEA Level 3"})
	performanceValues = append(performanceValues, models.NZPerformanceValue{Title: "Qualification", Name: "NCEA Level 1"})

	for _, qualification := range performanceValues {
		var performances = models.NZPerformance{Name: "Qualification"}
		performances.PerformanceValue = append(performances.PerformanceValue, qualification)
		err = db.Create(&performances).Error
		if err != nil {
			return err
		}

	}

	return nil
}

func addLit(db *gorm.DB) error {
	var err error
	var performanceValues []models.NZPerformanceValue
	performanceValues = append(performanceValues, models.NZPerformanceValue{Title: "Certificate", Name: "Literacy"})
	performanceValues = append(performanceValues, models.NZPerformanceValue{Title: "Certificate", Name: "Numeracy"})

	for _, qualification := range performanceValues {
		var performances = models.NZPerformance{Name: "Lit"}
		performances.PerformanceValue = append(performances.PerformanceValue, qualification)
		err = db.Create(&performances).Error
		if err != nil {
			return err
		}

	}

	return nil
}

func addYearLevel(db *gorm.DB) error {
	var err error

	var yealLevels []models.NZYearLevel
	yealLevels = append(yealLevels, models.NZYearLevel{Level: 11})
	yealLevels = append(yealLevels, models.NZYearLevel{Level: 12})
	yealLevels = append(yealLevels, models.NZYearLevel{Level: 13})
	for _, yearLevel := range yealLevels {
		err = db.Create(&yearLevel).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func addGender(db *gorm.DB) error {
	var err error

	var categories []models.NZCategory
	name := "Gender"
	categories = append(categories, models.NZCategory{Name: name, Value: "Male"})
	categories = append(categories, models.NZCategory{Name: name, Value: "Female"})
	categories = append(categories, models.NZCategory{Name: name, Value: "Unknown"})
	categories = append(categories, models.NZCategory{Name: name, Value: "All"})
	for _, category := range categories {
		err = db.Create(&category).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func addEthnicity(db *gorm.DB) error {
	var err error

	var categories []models.NZCategory
	name := "Ethnicity"
	categories = append(categories, models.NZCategory{Name: name, Value: "NZ Maori"})
	categories = append(categories, models.NZCategory{Name: name, Value: "Pasifika Peoples"})
	categories = append(categories, models.NZCategory{Name: name, Value: "NZ European"})
	categories = append(categories, models.NZCategory{Name: name, Value: "Asian"})
	categories = append(categories, models.NZCategory{Name: name, Value: "Other"})
	categories = append(categories, models.NZCategory{Name: name, Value: "All"})
	for _, category := range categories {
		err = db.Create(&category).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func addDecileBand(db *gorm.DB) error {
	var err error

	var categories []models.NZCategory
	name := "DecileBand"
	categories = append(categories, models.NZCategory{Name: name, Value: "Decile 8-10"})
	categories = append(categories, models.NZCategory{Name: name, Value: "Decile 1-3"})
	categories = append(categories, models.NZCategory{Name: name, Value: "Decile 4-7"})
	categories = append(categories, models.NZCategory{Name: name, Value: "Unknown"})
	categories = append(categories, models.NZCategory{Name: name, Value: "All"})
	for _, category := range categories {
		err = db.Create(&category).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func addNation(db *gorm.DB) error {
	var err error

	var categories []models.NZCategory
	name := "National"
	categories = append(categories, models.NZCategory{Name: name, Value: "National"})
	for _, category := range categories {
		err = db.Create(&category).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := ImportData()
	if err != nil {
		panic(err)
	}

}
