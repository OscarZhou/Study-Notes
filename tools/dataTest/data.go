package main

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/inflection"
)

var (
	tables map[string]interface{}
)

func init() {

	tables = make(map[string]interface{})

}

type Value struct {
	Ethnicity string
}

func FindData() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	var performanceValueMap map[string]bool

	performanceValueMap = make(map[string]bool)

	var keyword string
	keyword = "ethnicity"

	var values []Value
	for i, v := range tables {
		table := reflect.ValueOf(v).Elem()
		for idx := 0; idx < table.NumField(); idx++ {
			tableName := gorm.ToDBName(inflection.Singular(table.Type().Field(idx).Name))
			if tableName == keyword {
				templateString := "SELECT DISTINCT " + keyword + " FROM " + i + " GROUP BY " + keyword
				err = db.Raw(templateString).Scan(&values).Error
				if err != nil {
					return err
				}
				for _, vv := range values {
					key := vv.Ethnicity
					if _, ok := performanceValueMap[key]; !ok {
						performanceValueMap[key] = true
					}
				}
				break
			}

		}

	}

	for i, _ := range performanceValueMap {
		fmt.Println(i)
	}

	return nil
}
