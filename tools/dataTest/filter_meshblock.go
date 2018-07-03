package main

import (
	"Study-Notes/tools/dataTest/models"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func FilterMeshblock() error {
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

	// err = addAdjustGeoInfo(db)
	// if err != nil {
	// 	return err
	// }

	err = addDatazoneID2Meshblock(db)
	if err != nil {
		return err
	}

	err = reverseGeoPosition(db)
	if err != nil {
		return err
	}
	return nil
}

func reverseGeoPosition(db *gorm.DB) error {
	var meshblocks []models.Meshblock
	err := db.Find(&meshblocks).Error
	if err != nil {
		return err
	}

	tx := db.Begin()
	for _, v := range meshblocks {
		if v.WKT == "" {
			continue
		}
		fmt.Println("mb2018=", v.Mb2018)
		wkts := strings.Split(v.WKT, ",")
		for i, position := range wkts {
			pos := strings.Split(position, " ")
			pos[0], pos[1] = pos[1], pos[0]
			wkts[i] = strings.Join(pos, " ")
		}
		v.WKT = strings.Join(wkts, ",")
		err = tx.Table("meshblocks").Where("id = ?", v.ID).Updates(map[string]interface{}{"wkt": v.WKT}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func addAdjustGeoInfo(db *gorm.DB) error {
	var meshblocks []models.Meshblock
	err := db.Find(&meshblocks).Error
	if err != nil {
		return err
	}

	tx := db.Begin()
	for _, v := range meshblocks {
		fmt.Println("mb2018=", v.Mb2018)
		wkt := v.WKT
		wkt = strings.Replace(wkt, "MULTIPOLYGON (((", "", -1)
		wkt = strings.Replace(wkt, ")))", "", -1)
		v.WKT = wkt
		var meshblockOld models.MeshblockOld
		if tx.Where("mb2018_v1 = ?", v.Mb2018).Find(&meshblockOld).RecordNotFound() {
			continue
		}

		v.Latitude = meshblockOld.Latitude
		v.Longitude = meshblockOld.Longitude
		err = tx.Table("meshblocks").Where("id = ?", v.ID).Updates(map[string]interface{}{"wkt": v.WKT, "latitude": v.Latitude, "longitude": v.Longitude}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func addDatazoneID2Meshblock(db *gorm.DB) error {
	var meshblocks []models.Meshblock
	err := db.Find(&meshblocks).Error
	if err != nil {
		return err
	}

	tx := db.Begin()
	for _, v := range meshblocks {
		var meshblockMapping models.MeshblockMapping
		if tx.Where("mb2018 = ?", v.Mb2018).Find(&meshblockMapping).RecordNotFound() {
			if tx.Where("mb2013 = ?", v.Mb2018).Find(&meshblockMapping).RecordNotFound() {
				continue
			}
		}

		v.DatazoneID = meshblockMapping.DatazoneID
		err = tx.Table("meshblocks").Where("id = ?", v.ID).Updates(map[string]interface{}{"datazone_id": v.DatazoneID}).Error
		if err != nil {
			fmt.Println("id=", v.ID)
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
