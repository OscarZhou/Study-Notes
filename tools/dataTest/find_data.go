package main

import (
	"Study-Notes/tools/dataTest/models"
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
	tables["academic_decile_endorsements"] = &models.AcademicDecileEndorsement{}
	tables["academic_decile_lits"] = &models.AcademicDecileLit{}
	tables["academic_decile_qualifications"] = &models.AcademicDecileQualification{}

	tables["academic_ethnicity_endorsements"] = &models.AcademicEthnicityEndorsement{}
	tables["academic_ethnicity_lits"] = &models.AcademicEthnicityLit{}
	tables["academic_ethnicity_qualifications"] = &models.AcademicEthnicityQualification{}

	tables["academic_gender_endorsements"] = &models.AcademicGenderEndorsement{}
	tables["academic_gender_lits"] = &models.AcademicGenderLit{}
	tables["academic_gender_qualifications"] = &models.AcademicGenderQualification{}

	tables["ncea_academic_endorsements"] = &models.NCEAAcademicEndorsement{}
	tables["ncea_academic_lits"] = &models.NCEAAcademicLit{}
	tables["ncea_academic_qualifications"] = &models.NCEAAcademicQualification{}

	tables["ncea_school_academic_endorsements"] = &models.NCEASchoolAcademicEndorsement{}
	tables["ncea_school_academic_lits"] = &models.NCEASchoolAcademicLit{}
	tables["ncea_school_academic_qualifications"] = &models.NCEASchoolAcademicQualification{}

	tables["school_academic_ethnicity_endorsements"] = &models.SchoolAcademicEthnicityEndorsement{}
	tables["school_academic_ethnicity_lits"] = &models.SchoolAcademicEthnicityLit{}
	tables["school_academic_ethnicity_qualifications"] = &models.SchoolAcademicEthnicityQualification{}

	tables["school_academic_gender_endorsements"] = &models.SchoolAcademicGenderEndorsement{}
	tables["school_academic_gender_lits"] = &models.SchoolAcademicGenderLit{}
	tables["school_academic_gender_qualifications"] = &models.SchoolAcademicGenderQualification{}

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

	for i := range performanceValueMap {
		fmt.Println(i)
	}

	return nil
}
