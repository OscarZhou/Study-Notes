package main

import (
	"Study-Notes/tools/dataTest/models"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func AddSchoolURL() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	err = addSchoolURLSEO(db)
	if err != nil {
		return err
	}

	return nil
}

func addSchoolURLSEO(db *gorm.DB) error {
	var schools []models.School
	err := db.Find(&schools).Error
	if err != nil {
		return err
	}

	tx := db.Begin()
	for _, v := range schools {
		url := v.SchoolName
		url = strings.TrimSpace(url)
		url = strings.Replace(url, "'", "", -1)
		url = strings.Replace(url, " ", "-", -1)
		url = strings.Replace(url, " ", "-", -1)
		url = strings.Replace(url, "(", "", -1)
		url = strings.Replace(url, ")", "", -1)
		url = strings.Replace(url, "（", "", -1)
		url = strings.Replace(url, "）", "", -1)
		url = strings.Replace(url, "&", "-and-", -1)
		url = strings.Replace(url, "|", "-or-", -1)
		url = strings.Replace(url, ".", "", -1)
		url = strings.Replace(url, "/", "", -1)
		url = strings.Replace(url, ":", "-", -1)
		url = strings.Replace(url, "--", "-", -1)
		url = strings.ToLower(url)
		url = url + "-nz"
		v.URL = url
		if err := tx.Save(&v).Error; err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}
