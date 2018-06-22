package main

import (
	"Study-Notes/tools/dataTest/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func ImportSchoolRollCategory() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	db.DropTable(
		models.SchoolRollCategory{},
	)

	db.AutoMigrate(
		models.SchoolRollCategory{},
	)

	// db.LogMode(true)

	// Ethnicity Rolls
	err = loadEthnicityRoll(db)
	if err != nil {
		return err
	}
	// Gender Rolls
	err = loadGenderRoll(db)
	if err != nil {
		return err
	}

	return nil
}

func loadEthnicityRoll(db *gorm.DB) error {

	var ethnicities []models.NZCategory
	err := db.Where("name = ? and parent_id = ?", "Ethnicity", 0).Find(&ethnicities).Error
	if err != nil {
		return err
	}

	var schoolRolls []models.SchoolRoll
	err = db.Find(&schoolRolls).Error
	if err != nil {
		return err
	}

	var schoolRollCategories []models.SchoolRollCategory
	for _, v := range schoolRolls {
		for _, ethnicity := range ethnicities {
			switch ethnicity.Value {
			case "European":
				schoolRollCategories = append(schoolRollCategories, models.SchoolRollCategory{
					SchoolID:   v.SchoolID,
					Year:       v.Year,
					Roll:       v.EuropeanPakehaRoll,
					CategoryID: ethnicity.ID,
				})
			case "Maori":
				schoolRollCategories = append(schoolRollCategories, models.SchoolRollCategory{
					SchoolID:   v.SchoolID,
					Year:       v.Year,
					Roll:       v.MaoriRoll,
					CategoryID: ethnicity.ID,
				})
			case "Asian":
				schoolRollCategories = append(schoolRollCategories, models.SchoolRollCategory{
					SchoolID:   v.SchoolID,
					Year:       v.Year,
					Roll:       v.AsianRoll,
					CategoryID: ethnicity.ID,
				})
			case "MELAA":
				schoolRollCategories = append(schoolRollCategories, models.SchoolRollCategory{
					SchoolID:   v.SchoolID,
					Year:       v.Year,
					Roll:       v.MelaaRoll,
					CategoryID: ethnicity.ID,
				})
			case "Pacifika Peoples":
				schoolRollCategories = append(schoolRollCategories, models.SchoolRollCategory{
					SchoolID:   v.SchoolID,
					Year:       v.Year,
					Roll:       v.PasifikaRoll,
					CategoryID: ethnicity.ID,
				})
			case "All":
				schoolRollCategories = append(schoolRollCategories, models.SchoolRollCategory{
					SchoolID:   v.SchoolID,
					Year:       v.Year,
					Roll:       v.TotalRoll,
					CategoryID: ethnicity.ID,
				})
			}
		}
	}
	tx := db.Begin()
	for _, v := range schoolRollCategories {
		err = tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	fmt.Printf("ethnicity: %d", len(schoolRollCategories))
	return nil
}

func loadGenderRoll(db *gorm.DB) error {

	var genders []models.NZCategory
	err := db.Where("name = ? and parent_id = ?", "Gender", 0).Find(&genders).Error
	if err != nil {
		return err
	}

	var schoolRolls []models.SchoolRoll
	err = db.Find(&schoolRolls).Error
	if err != nil {
		return err
	}

	var schoolRollCategories []models.SchoolRollCategory
	for _, v := range schoolRolls {
		for _, gender := range genders {
			switch gender.Value {
			case "Male":
				schoolRollCategories = append(schoolRollCategories, models.SchoolRollCategory{
					SchoolID:   v.SchoolID,
					Year:       v.Year,
					Roll:       v.MaleRoll,
					CategoryID: gender.ID,
				})
			case "Female":
				schoolRollCategories = append(schoolRollCategories, models.SchoolRollCategory{
					SchoolID:   v.SchoolID,
					Year:       v.Year,
					Roll:       v.FemaleRoll,
					CategoryID: gender.ID,
				})
			case "All":
				schoolRollCategories = append(schoolRollCategories, models.SchoolRollCategory{
					SchoolID:   v.SchoolID,
					Year:       v.Year,
					Roll:       v.TotalRoll,
					CategoryID: gender.ID,
				})
			}
		}
	}

	tx := db.Begin()
	for _, v := range schoolRollCategories {
		err = tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	fmt.Printf("gender: %d", len(schoolRollCategories))
	return nil
}
