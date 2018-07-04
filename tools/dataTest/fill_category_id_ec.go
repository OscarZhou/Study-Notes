package main

import (
	"Study-Notes/tools/dataTest/models"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func FillCategoryID2ExtraCurriculum() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	// db.LogMode(true)

	// invalidEC, err := fillExtraCurriculumActivityType(db)
	// if err != nil {
	// 	return err
	// }

	// err = export2Csv(invalidEC)
	// if err != nil {
	// 	return err
	// }
	invalidEC, err := loadUnknownFields(db)
	if err != nil {
		return err
	}
	err = export2Csv1(invalidEC)
	if err != nil {
		return err
	}

	return nil
}

func fillExtraCurriculumActivityType(db *gorm.DB) ([]models.ExtraCurriculum, error) {
	var (
		extraCurriculums []models.ExtraCurriculum
		invalidEC        []models.ExtraCurriculum
	)
	err := db.Find(&extraCurriculums).Error
	if err != nil {
		return nil, err
	}

	count1, count2, count3 := 0, 0, 0
	tx := db.Begin()
	for i, v := range extraCurriculums {
		if v.ActivityType == "" {
			count1++
			invalidEC = append(invalidEC, v)
			continue
		}

		v.ActivityType = strings.Replace(v.ActivityType, "\n", "", -1)
		// v.ActivityType = strings.Replace(v.ActivityType, "&", " ", -1)
		v.ActivityType = strings.Replace(v.ActivityType, "  ", " ", -1)
		activityTypes := strings.Split(v.ActivityType, ",")
		for i, t := range activityTypes {

			activityTypes[i] = strings.TrimSpace(t)
			if activityTypes[i] == "Athletics" {
				activityTypes[i] = "Athletic"
			}
		}

		for _, actv := range activityTypes {
			activityType := "%" + strings.TrimSpace(actv) + "%"
			var categories []models.NZCategory
			if tx.Where("name = ? and parent_id <> ? and value ilike ?", "ExtraCurriculum", 0, activityType).
				Find(&categories).RecordNotFound() {
				fmt.Printf("ec id = %d, name = %s, type = %s, category = %v\n", v.ID, v.Name, activityType, categories)
				continue
			}
			extraCurriculums[i].Categories = append(extraCurriculums[i].Categories, categories...)
		}

		if len(extraCurriculums[i].Categories) == 0 {
			// fmt.Printf("ec id = %d, name = %s, categories = %v, len = %d\n", v.ID, v.Name, categories, len(categories))
			count3++
			invalidEC = append(invalidEC, v)
			continue
		}

		err = tx.Save(&extraCurriculums[i]).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		fmt.Println(i)
	}
	tx.Commit()

	fmt.Println("count1=", count1)
	fmt.Println("count2=", count2)
	fmt.Println("count3=", count3)
	// for _, v := range invalidEC {
	// 	fmt.Printf("ec id = %d, name = %s, activity type = %s\n", v.ID, v.Name, v.ActivityType)
	// }
	fmt.Println("invalid number is ", len(invalidEC))
	return invalidEC, nil
}

func export2Csv(extraCurr []models.ExtraCurriculum) error {
	file, err := os.Create("unknownCategory.csv")
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	var data [][]string
	// add title
	var title []string
	title, err = getFieldName(&models.ExtraCurriculum{})
	data = append(data, title)

	// add content
	// allex := reflect.ValueOf(extraCurr)
	// if allex.IsValid() {
	// 	if allex.Kind() == reflect.Slice {
	// 		for i := 0; i < allex.Len(); i++ {
	// 			if allex.Index(i).Kind() == reflect.Struct {
	// 				for j := 0; j < allex.Index(i).NumField(); j++ {

	// 					var row []string
	// 					row = append(row, allex.Index(i).Field(j).String())
	// 					data = append(data, row)
	// 					fmt.Println(row)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	fmt.Println(title)
	for _, v := range extraCurr {
		id := strconv.FormatUint(uint64(v.ID), 10)
		lat := strconv.FormatFloat(v.Latitude, 'f', 1, 64)
		lon := strconv.FormatFloat(v.Longitude, 'f', 1, 64)
		row := []string{id, "", "", "", v.Name, v.PhoneNumber, v.Email, v.Website, v.ContactPerson,
			v.Street, v.Suburb, v.City, v.Country, v.PostalAddress1, v.PostalAddress2, v.PostalAddress3, v.PostalCode, lat, lon,
			v.ActivityType, v.BookingLink, v.OpeningHours, v.Monday, v.Tuesday, v.Wednesday, v.Thursday, v.Friday, v.Saturday, v.Sunday, v.Facebook,
			v.Twitter}
		data = append(data, row)
	}
	for _, v := range data {
		err := writer.Write(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func getFieldName(model interface{}) ([]string, error) {
	var (
		fields []string
	)
	tableStruct := reflect.ValueOf(model)
	if tableStruct.Kind() == reflect.Ptr {
		tableStruct = tableStruct.Elem()
	}

	if !tableStruct.IsValid() {
		return nil, errors.New("model is invalid")
	}
	switch tableStruct.Kind() {
	case reflect.Struct:
		for i := 0; i < tableStruct.NumField(); i++ {
			name := tableStruct.Type().Field(i).Name
			if strings.Contains(name, "At") {
				fields = append(fields, name)
				continue
			} else if name == "Categories" {
				continue
			} else {
				// fmt.Println(tableStruct.Field(i).Interface())
				subFields, err := getFieldName(tableStruct.Field(i).Interface())
				if err != nil {
					return nil, err
				}
				if len(subFields) == 0 {

					fields = append(fields, name)
				} else {
					fields = append(fields, subFields...)
				}
			}

		}
	case reflect.Slice:
		return fields, nil
	default:
		return fields, nil
	}
	return fields, nil

}

func loadUnknownFields(db *gorm.DB) ([][]string, error) {
	var (
		extraCurriculums []models.ExtraCurriculum
		invalidEC        [][]string
	)
	err := db.Find(&extraCurriculums).Error
	if err != nil {
		return nil, err
	}

	invalidEC = append(invalidEC, []string{"ID", "Activity Type"})
	count1, count2, count3 := 0, 0, 0
	tx := db.Begin()
	for _, v := range extraCurriculums {
		if v.ActivityType == "" {
			count1++
			continue
		}

		v.ActivityType = strings.Replace(v.ActivityType, "\n", "", -1)
		// v.ActivityType = strings.Replace(v.ActivityType, "&", " ", -1)
		v.ActivityType = strings.Replace(v.ActivityType, "  ", " ", -1)
		activityTypes := strings.Split(v.ActivityType, ",")
		for i, t := range activityTypes {
			activityTypes[i] = strings.TrimSpace(t)
			if activityTypes[i] == "Athletics" {
				activityTypes[i] = "Athletic"
			}
		}

		for _, actv := range activityTypes {
			activityType := "%" + strings.TrimSpace(actv) + "%"
			var categories []models.NZCategory
			if tx.Where("name = ? and parent_id <> ? and value ilike ?", "ExtraCurriculum", 0, activityType).
				Find(&categories).RecordNotFound() {
				fmt.Printf("ec id = %d, name = %s, type = %s, category = %v\n", v.ID, v.Name, activityType, categories)
				id := strconv.FormatUint(uint64(v.ID), 10)
				invalidEC = append(invalidEC, []string{id, v.ActivityType[1 : len(v.ActivityType)-1]})
				continue
			}

			if len(categories) == 0 {
				fmt.Printf("ec id = %d, name = %s, type = %s, category = %v\n", v.ID, v.Name, activityType, categories)
				id := strconv.FormatUint(uint64(v.ID), 10)
				invalidEC = append(invalidEC, []string{id, activityType[1 : len(activityType)-1]})
				count3++
			}
		}

		// fmt.Println(i)
	}
	tx.Commit()

	fmt.Println("count1=", count1)
	fmt.Println("count2=", count2)
	fmt.Println("count3=", count3)
	// for _, v := range invalidEC {
	// 	fmt.Printf("ec id = %d, name = %s, activity type = %s\n", v.ID, v.Name, v.ActivityType)
	// }
	fmt.Println("invalid number is ", len(invalidEC))
	return invalidEC, nil
}

func export2Csv1(extraCurr [][]string) error {
	file, err := os.Create("unknownCategory.csv")
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, v := range extraCurr {
		err := writer.Write(v)
		if err != nil {
			return err
		}
	}

	return nil
}
