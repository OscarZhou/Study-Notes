package main

import (
	"Study-Notes/tools/dataTest/models"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func AddUnknownExtraCurriculumCategory() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	db.LogMode(true)
	// db.DropTable(
	// 	models.UnknownECCategory{},
	// )

	// db.AutoMigrate(
	// 	models.UnknownECCategory{},
	// )

	err = addUnknownECCategory(db)
	if err != nil {
		return err
	}

	return nil
}

func addUnknownECCategory(db *gorm.DB) error {
	var unknownCategories []models.UnknownECCategory
	if err := db.Find(&unknownCategories).Error; err != nil {
		return err
	}

	tx := db.Begin()
	for _, v := range unknownCategories {
		var ec models.ExtraCurriculum
		if err := tx.Where("id = ?", v.ECID).Find(&ec).Error; err != nil {
			tx.Rollback()
			return err
		}

		var category models.NZCategory
		if err := tx.Where("name = ? and value = ?", "ExtraCurriculum", strings.TrimSpace(v.Activity)).Find(&category).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Model(&ec).Association("Categories").Clear().Error; err != nil {
			tx.Rollback()
			return err
		}

		ec.Categories = append(ec.Categories, category)

		if err := tx.Save(&ec).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
