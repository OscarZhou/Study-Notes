package main

import (
	"Study-Notes/tools/dataTest/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func AddExtraCurriculumCategory() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	// db.DropTable(
	// 	models.ExtraCCCategory{},
	// )

	// db.AutoMigrate(
	// 	models.ExtraCCCategory{},
	// )

	err = AddECLevel1Type2Category(db)
	if err != nil {
		return err
	}

	err = AddECLevel2Type2Category(db)
	if err != nil {
		return err
	}
	return nil
}

func AddECLevel2Type2Category(db *gorm.DB) error {
	var extraCurriculumCategories []models.ExtraCCCategory
	err := db.Find(&extraCurriculumCategories).Error
	if err != nil {
		return err
	}
	tx := db.Begin()
	for _, v := range extraCurriculumCategories {
		var parentCategory models.NZCategory
		err = tx.Where("name = ? and value = ? and parent_id = ?", "ExtraCurriculum", v.Type, 0).Find(&parentCategory).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		var category models.NZCategory
		category.Name = parentCategory.Name
		category.Value = v.Name
		category.ParentID = parentCategory.ID
		err = tx.Create(&category).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func AddECLevel1Type2Category(db *gorm.DB) error {
	var extraCurriculumCategories []models.ExtraCCCategory
	err := db.Find(&extraCurriculumCategories).Error
	if err != nil {
		return err
	}

	var ecCategories = make(map[string]bool)
	for _, v := range extraCurriculumCategories {
		_, exist := ecCategories[v.Type]
		if !exist {
			ecCategories[v.Type] = true
		}
	}

	tx := db.Begin()
	for k, v := range ecCategories {
		fmt.Println(k, v)
		var dupeCategory models.NZCategory
		dupeCategory.Name = "ExtraCurriculum"
		dupeCategory.Value = k
		dupeCategory.ParentID = 0
		err = tx.Create(&dupeCategory).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
