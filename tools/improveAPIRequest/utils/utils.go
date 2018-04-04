package utils

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

func ConvertJSON2String(object interface{}) (string, error) {
	b, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GetDBTableNameByModel(db *gorm.DB, model interface{}) string {
	dbTableName := db.NewScope(model).TableName()
	return dbTableName
}
