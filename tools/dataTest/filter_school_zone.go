package main

import (
	"Study-Notes/tools/dataTest/models"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func FilterSchoolZone() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	// db.DropTable(
	// 	models.SchoolZone{},
	// )

	// db.AutoMigrate(
	// 	models.SchoolZone{},
	// )

	err = groupGeoPosition(db)
	if err != nil {
		return err
	}

	// err = removeAllComma(db)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func groupGeoPosition(db *gorm.DB) error {
	var schoolZones []models.SchoolZone
	err := db.Preload("Surfaces").Find(&schoolZones).Error
	if err != nil {
		return err
	}

	tx := db.Begin()
	for _, v := range schoolZones {
		for si, surface := range v.Surfaces {
			var newList []string
			posLists := strings.Split(surface.PosList, " ")
			for i, pv := range posLists {
				if (i+1)%2 == 0 {
					geo := posLists[i-1] + " " + pv
					newList = append(newList, geo)
				}
			}
			v.Surfaces[si].PosList = strings.Join(newList, ",")
			err = tx.Table("school_zone_surfaces").Where("school_zone_id = ?", v.Surfaces[si].SchoolZoneID).Updates(map[string]interface{}{"pos_list": v.Surfaces[si].PosList}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			fmt.Println("SchoolZoneID = ", surface.SchoolZoneID)
		}
	}

	tx.Commit()
	return nil
}

func removeAllComma(db *gorm.DB) error {
	var schoolZones []models.SchoolZone
	err := db.Preload("Surfaces").Find(&schoolZones).Error
	if err != nil {
		return err
	}

	tx := db.Begin()
	for _, v := range schoolZones {
		for si, surface := range v.Surfaces {
			v.Surfaces[si].PosList = strings.Replace(surface.PosList, ",", " ", -1)
			err = tx.Table("school_zone_surfaces").Where("school_zone_id = ?", v.Surfaces[si].SchoolZoneID).Updates(map[string]interface{}{"pos_list": v.Surfaces[si].PosList}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			fmt.Println("SchoolZoneID = ", surface.SchoolZoneID)
		}
	}
	tx.Commit()
	return nil
}
