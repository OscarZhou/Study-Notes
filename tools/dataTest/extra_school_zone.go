package main

import (
	"Study-Notes/tools/dataTest/models"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ExtraSchoolZone() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	if err := extraSchoolZonePos(db); err != nil {
		return err
	}

	return nil
}

func extraSchoolZonePos(db *gorm.DB) error {
	var (
		schoolZoneSurface []models.SchoolZoneSurface
		schoolZonePos     string
	)
	if err := db.Order("school_zone_id asc").Find(&schoolZoneSurface).Error; err != nil {
		return err
	}

	for _, v := range schoolZoneSurface {
		posList := strings.Replace(v.PosList, ",", "\r\n", -1)
		posList += "\r\n0 0\r\n"
		schoolZonePos += posList

	}

	if err := createJSONFile("school_zone_pos.txt", schoolZonePos); err != nil {
		return err
	}
	// fmt.Println(schoolZonePos)
	return nil
}

func createJSONFile(filename, content string) error {
	var (
		file *os.File
	)

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}
