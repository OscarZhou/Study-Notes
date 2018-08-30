package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/dongri/phonenumber"
)

func main() {
	if err := ExportGlobalCountryBasicData(); err != nil {
		panic(err)
	}
}

type Country struct {
	Name        string
	Alpha2      string `gorm:"size:2"`
	Alpha3      string `gorm:"size:3"`
	CallingCode string
}

func ExportGlobalCountryBasicData() error {
	iso3166 := phonenumber.GetISO3166()
	rows := make([][]string, len(iso3166))
	for i, v := range iso3166 {
		row := []string{v.CountryName, v.Alpha2, v.Alpha3, v.CountryCode}
		rows[i] = append(rows[i], row...)
	}

	if err := Export2CSV("global_country.csv", &Country{}, rows...); err != nil {
		return err
	}
	return nil
}

func Export2CSV(filename string, model interface{}, rows ...[]string) error {
	var (
		csvData [][]string
		title   []string
	)
	if filepath.Ext(filename) != ".csv" {
		filename += ".csv"
	}

	file, err := os.Create("./" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if title, err = GetStructFieldNames(model); err != nil {
		return err
	}

	length := len(title)
	csvData = append(csvData, title)
	for _, row := range rows {
		if len(row) != length {
			return fmt.Errorf("the length between title and data is inconsistent. title:%d, data:%d", length, len(row))
		}
		csvData = append(csvData, row)
	}

	for _, v := range csvData {
		if err = writer.Write(v); err != nil {
			return err
		}
	}

	return nil
}

func GetStructFieldNames(model interface{}) ([]string, error) {
	var fields []string
	modelStruct := reflect.ValueOf(model)
	if modelStruct.Kind() == reflect.Ptr {
		modelStruct = modelStruct.Elem()
	}

	if !modelStruct.IsValid() {
		return nil, errors.New("model is invalid")
	}

	switch modelStruct.Kind() {
	case reflect.Struct:
		for i := 0; i < modelStruct.NumField(); i++ {
			name := modelStruct.Type().Field(i).Name
			typ := modelStruct.Field(i).Type().String()
			if typ == "time.Time" || typ == "*time.Time" {
				fields = append(fields, name)
				continue
			}

			if modelStruct.Field(i).Kind() == reflect.Slice ||
				modelStruct.Field(i).Kind() == reflect.Struct {
				continue
			}
			fields = append(fields, name)
		}
	case reflect.Slice:
		return fields, nil
	default:
		return fields, nil
	}
	return fields, nil
}
