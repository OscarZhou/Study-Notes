package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"osm_library/models/info"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	data, err := ioutil.ReadFile("citydata.json")
	if err != nil {
		fmt.Println("%v", err)
		return
	}

	var china []info.Province
	err = json.Unmarshal(data, &china)
	if err != nil {
		fmt.Println("error:", err)
	}
	var (
		db      *gorm.DB
		dialect string
		dbInfo  string
	)
	dialect = "postgres"
	dbInfo = "host=192.168.164.129 port=5432 user=osmapi dbname=osmapi password=osm123"

	db, err = gorm.Open(dialect, dbInfo)
	if err != nil {
		log.Println("failed to connect database!")
		return
	}

	for _, v := range china {
		db.Create(&v)
	}

	fmt.Println(china)

}
