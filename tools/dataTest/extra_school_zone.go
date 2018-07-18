package main

import (
	"Study-Notes/tools/dataTest/models"
	"bufio"
	"fmt"
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

	// if err := extraSchoolZonePos(db); err != nil {
	// 	return err
	// }

	if err := updateSchoolZonePos(db); err != nil {
		return err
	}

	return nil
}

func updateSchoolZonePos(db *gorm.DB) error {
	file, err := os.Open("processed_zone.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 2*1024*1024)
	tx := db.Begin()
	i := 0
	for scanner.Scan() {
		i++
		t := scanner.Text()
		if err := tx.Table("school_zone_surfaces").
			Where("school_zone_id = ?", i).
			Updates(map[string]interface{}{"pos_list": t[0 : len(t)-1]}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	if err := scanner.Err(); err != nil {
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
		posList := strings.Replace(v.PosList, ",", "\n", -1)
		posList += "\n0 0\n"
		schoolZonePos += posList

	}

	if err := createJSONFile("school_zone_pos.txt", schoolZonePos); err != nil {
		return err
	}
	fmt.Println(schoolZonePos)
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
