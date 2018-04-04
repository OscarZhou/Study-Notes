package main

import (
	"Study-Notes/tools/improveAPIRequest/configs"
	"Study-Notes/tools/improveAPIRequest/models"

	"github.com/jinzhu/gorm"
)

//InitDB initial gorm db
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(configs.LocalConfig.DbEngin,
		configs.LocalConfig.DbConn)
	if err != nil {
		return nil, err
	}
	if configs.LocalConfig.DropTables {
		db.DropTableIfExists()
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB,
		defaultTableName string) string {
		return defaultTableName
	}

	if configs.LocalConfig.AutoMigrate {
		if err := db.AutoMigrate(
			&models.ExchangeRate{},
			&models.CurrencyRelation{},
		).Error; err != nil {
			return nil, err
		}
	}

	db.LogMode(false)

	return db, nil
}
