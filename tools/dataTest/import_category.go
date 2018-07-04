package main

import (
	"Study-Notes/tools/dataTest/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func ImportCategory() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	db.DropTable(
		models.NZCategory{},
	)

	db.AutoMigrate(
		models.NZCategory{},
	)

	db.LogMode(true)

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

	var category models.NZCategory
	// European
	err = db.Where("name = ? and value = ?", "Ethnicity", "European").Find(&category).Error
	if err != nil {
		return err
	}

	categories := []models.NZCategory{
		models.NZCategory{
			Value:    "European not further defined",
			Name:     category.Name,
			ParentID: category.ID,
		},
		models.NZCategory{
			Value:    "NZ European",
			Name:     category.Name,
			ParentID: category.ID,
		},
		models.NZCategory{
			Value:    "Other European",
			Name:     category.Name,
			ParentID: category.ID,
		},
	}
	// Maori
	category = models.NZCategory{}
	err = db.Where("name = ? and value = ?", "Ethnicity", "Maori").Find(&category).Error
	if err != nil {
		return err
	}

	categories = append(categories, models.NZCategory{Value: "NZ Maori", Name: category.Name, ParentID: category.ID})
	// Pacific People
	category = models.NZCategory{}
	err = db.Where("name = ? and value = ?", "Ethnicity", "Pasifika Peoples").Find(&category).Error
	if err != nil {
		return err
	}

	categories = append(categories, models.NZCategory{Value: "Pacific Island not further defined", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Samoan", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Cook Island Maori", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Tongan", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Niuean", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Tokelauan", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Fijian", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Other Pacific Island", Name: category.Name, ParentID: category.ID})

	// Asian
	category = models.NZCategory{}
	err = db.Where("name = ? and value = ?", "Ethnicity", "Asian").Find(&category).Error
	if err != nil {
		return err
	}

	categories = append(categories, models.NZCategory{Value: "Asian not further defined", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Southeast Asian", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Chinese", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Indian", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Other Asian", Name: category.Name, ParentID: category.ID})

	// MELAA

	tx := db.Begin()
	category = models.NZCategory{Name: "Ethnicity", Value: "MELAA", ParentID: 0}
	err = tx.Create(&category).Find(&category).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	categories = append(categories, models.NZCategory{Value: "Middle Eastern", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Latin American / Hispanic", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "African", Name: category.Name, ParentID: category.ID})
	categories = append(categories, models.NZCategory{Value: "Other", Name: category.Name, ParentID: category.ID})

	for _, v := range categories {
		err = tx.Create(&v).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	fmt.Printf("new insert: %d\n", len(categories))

	return nil
}
