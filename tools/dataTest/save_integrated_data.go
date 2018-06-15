package main

import (
	"Study-Notes/tools/dataTest/models"
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SaveIntegratedData() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	db.DropTable(
		models.NZAcademic{},
	)

	db.AutoMigrate(
		models.NZAcademic{},
	)

	// var academics []models.NZAcademic
	// for _, v := range tables {
	// 	tableType := reflect.ValueOf(v).Elem().Type()
	// 	fmt.Println(tableType.Name())

	// 	arrayTableType := reflect.SliceOf(tableType)
	// 	fmt.Println(arrayTableType)

	// 	slice := reflect.MakeSlice(arrayTableType, 100, 100)
	// 	fmt.Println(slice.Kind())
	// 	fmt.Println(slice.Len())
	// 	fmt.Println(slice.Cap())

	// 	ptr := reflect.New(slice.Type())
	// 	fmt.Println(ptr.Kind())
	// 	fmt.Println(ptr.Elem().Kind())
	// 	ptr.Elem().Set(slice)

	// 	fmt.Println(ptr)

	// 	fmt.Println(ptr.CanAddr())
	// 	if err := db.Limit(100).Find(ptr.Interface()).Error; err != nil {
	// 		return err
	// 	}

	// }
	// return nil

	// fmt.Println(academics)

	//1
	var academics []models.NZAcademic
	if err := loadAcademicDecileEndorsement(db, &academics); err != nil {
		return err
	}

	//2
	if err := loadAcademicDecileLit(db, &academics); err != nil {
		return err
	}

	//3
	if err := loadAcademicDecileQualification(db, &academics); err != nil {
		return err
	}
	//4
	if err := loadAcademicEthnicityEndorsement(db, &academics); err != nil {
		return err
	}

	//5
	if err := loadAcademicEthnicityLit(db, &academics); err != nil {
		return err
	}

	//6
	if err := loadAcademicEthnicityQualification(db, &academics); err != nil {
		return err
	}

	//7
	if err := loadAcademicGenderEndorsement(db, &academics); err != nil {
		return err
	}

	//8
	if err := loadAcademicGenderLit(db, &academics); err != nil {
		return err
	}

	//9
	if err := loadAcademicGenderQualification(db, &academics); err != nil {
		return err
	}

	//10
	if err := loadNCEAAcademicEndorsement(db, &academics); err != nil {
		return err
	}

	//11
	if err := loadNCEAAcademicLit(db, &academics); err != nil {
		return err
	}

	//12
	if err := loadNCEAAcademicQualification(db, &academics); err != nil {
		return err
	}

	for _, v := range academics {

		err := db.Create(&v).Error
		if err != nil {
			fmt.Println(v.ID)
			return err
		}
	}
	return nil

}

func loadAcademicDecileEndorsement(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.AcademicDecileEndorsement
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels        []models.NZYearLevel
			categories        []models.NZCategory
			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// DecileBand
		err = db.Where("name = ?", "DecileBand").Find(&categories).Error
		if err != nil {
			return err
		}
		for _, category := range categories {
			// fmt.Println(category.Value)
			// fmt.Println(v.DecileBand)

			if category.Value == v.DecileBand {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}

		err = db.Where("name = ?", "Endorsement").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and (ipv.title = ? OR ipv.title = ?)`, "Endorsement", "Qualification", "Endorsement").
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadAcademicDecileLit(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.AcademicDecileLit
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels        []models.NZYearLevel
			categories        []models.NZCategory
			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// DecileBand, Gender, Ethnicity
		err = db.Where("name = ?", "DecileBand").Find(&categories).Error
		if err != nil {
			return err
		}
		for _, category := range categories {
			// fmt.Println(category.Value)
			// fmt.Println(v.DecileBand)

			if category.Value == v.DecileBand {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}

		// Endorsement, Lit, Qualifaction
		err = db.Where("name = ?", "Lit").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and ipv.title in (?)`, "Certificate", []string{"Certificate"}).
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadAcademicDecileQualification(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.AcademicDecileQualification
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels        []models.NZYearLevel
			categories        []models.NZCategory
			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// DecileBand, Gender, Ethnicity
		err = db.Where("name = ?", "DecileBand").Find(&categories).Error
		if err != nil {
			return err
		}
		for _, category := range categories {
			// fmt.Println(category.Value)
			// fmt.Println(v.DecileBand)

			if category.Value == v.DecileBand {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}

		// Endorsement, Lit, Qualifaction
		err = db.Where("name = ?", "Qualification").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and ipv.title in (?)`, "Qualification", []string{"Qualification"}).
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadAcademicEthnicityEndorsement(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.AcademicEthnicityEndorsement
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels        []models.NZYearLevel
			categories        []models.NZCategory
			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// DecileBand, Gender, Ethnicity
		err = db.Where("name = ?", "Ethnicity").Find(&categories).Error
		if err != nil {
			return err
		}
		for _, category := range categories {
			// fmt.Println(category.Value)
			// fmt.Println(v.DecileBand)

			if category.Value == v.Ethnicity {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}

		// Endorsement, Lit, Qualifaction
		err = db.Where("name = ?", "Endorsement").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and ipv.title in (?)`, "Endorsement", []string{"Endorsement", "Qualification"}).
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadAcademicEthnicityLit(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.AcademicEthnicityLit
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels        []models.NZYearLevel
			categories        []models.NZCategory
			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// DecileBand, Gender, Ethnicity
		err = db.Where("name = ?", "Ethnicity").Find(&categories).Error
		if err != nil {
			return err
		}
		for _, category := range categories {
			// fmt.Println(category.Value)
			// fmt.Println(v.DecileBand)

			if category.Value == v.Ethnicity {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}

		// Endorsement, Lit, Qualifaction
		err = db.Where("name = ?", "Lit").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and ipv.title in (?)`, "Certificate", []string{"Certificate"}).
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadAcademicEthnicityQualification(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.AcademicEthnicityQualification
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels        []models.NZYearLevel
			categories        []models.NZCategory
			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// DecileBand, Gender, Ethnicity
		err = db.Where("name = ?", "Ethnicity").Find(&categories).Error
		if err != nil {
			return err
		}
		for _, category := range categories {
			// fmt.Println(category.Value)
			// fmt.Println(v.DecileBand)

			if category.Value == v.Ethnicity {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}

		// Endorsement, Lit, Qualifaction
		err = db.Where("name = ?", "Qualification").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and ipv.title in (?)`, "Qualification", []string{"Qualification"}).
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadAcademicGenderEndorsement(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.AcademicGenderEndorsement
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels        []models.NZYearLevel
			categories        []models.NZCategory
			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// DecileBand, Gender, Ethnicity
		err = db.Where("name = ?", "Gender").Find(&categories).Error
		if err != nil {
			return err
		}
		for _, category := range categories {
			// fmt.Println(category.Value)
			// fmt.Println(v.DecileBand)

			if category.Value == v.Gender {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}

		// Endorsement, Lit, Qualifaction
		err = db.Where("name = ?", "Endorsement").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and ipv.title in (?)`, "Endorsement", []string{"Endorsement", "Qualification"}).
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadAcademicGenderLit(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.AcademicGenderLit
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels        []models.NZYearLevel
			categories        []models.NZCategory
			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// DecileBand, Gender, Ethnicity
		err = db.Where("name = ?", "Gender").Find(&categories).Error
		if err != nil {
			return err
		}
		for _, category := range categories {
			// fmt.Println(category.Value)
			// fmt.Println(v.DecileBand)

			if category.Value == v.Gender {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}

		// Endorsement, Lit, Qualifaction
		err = db.Where("name = ?", "Lit").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and ipv.title in (?)`, "Certificate", []string{"Certificate"}).
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadAcademicGenderQualification(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.AcademicGenderQualification
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels        []models.NZYearLevel
			categories        []models.NZCategory
			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// DecileBand, Gender, Ethnicity
		err = db.Where("name = ?", "Gender").Find(&categories).Error
		if err != nil {
			return err
		}
		for _, category := range categories {
			// fmt.Println(category.Value)
			// fmt.Println(v.DecileBand)

			if category.Value == v.Gender {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}

		// Endorsement, Lit, Qualifaction
		err = db.Where("name = ?", "Qualification").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and ipv.title in (?)`, "Qualification", []string{"Qualification"}).
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadNCEAAcademicEndorsement(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.NCEAAcademicEndorsement
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var category models.NZCategory
	// National
	err = db.Where("name = ?", "National").Find(&category).Error
	if err != nil {
		return err
	}
	for _, v := range entities {
		var (
			yealLevels []models.NZYearLevel

			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate
		academic.CategoryID = category.ID

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		err = db.Where("name = ?", "Endorsement").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and (ipv.title = ? OR ipv.title = ?)`, "Endorsement", "Qualification", "Endorsement").
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadNCEAAcademicLit(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.NCEAAcademicLit
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var category models.NZCategory
	// National
	err = db.Where("name = ?", "National").Find(&category).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels []models.NZYearLevel

			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate
		academic.CategoryID = category.ID

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// Endorsement, Lit, Qualifaction
		err = db.Where("name = ?", "Lit").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and ipv.title in (?)`, "Certificate", []string{"Certificate"}).
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}

func loadNCEAAcademicQualification(db *gorm.DB, academics *[]models.NZAcademic) error {
	var (
		entities []models.NCEAAcademicQualification
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var category models.NZCategory
	// National
	err = db.Where("name = ?", "National").Find(&category).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			yealLevels        []models.NZYearLevel
			category          models.NZCategory
			performance       models.NZPerformance
			academic          models.NZAcademic
			performanceValues []models.NZPerformanceValue
		)

		academic.Year = v.Year
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate
		academic.ID = category.ID

		err = db.Find(&yealLevels).Error
		if err != nil {
			return err
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		// Endorsement, Lit, Qualifaction
		err = db.Where("name = ?", "Qualification").Find(&performance).Error
		if err != nil {
			return err
		}

		err = db.
			Raw(`select ipv.* from info_nz_performances ip
			join info_nz_performance_values ipv on ipv.performance_id = ip.id
			where ip.name = ? and ipv.title in (?)`, "Qualification", []string{"Qualification"}).
			Find(&performanceValues).Error
		if err != nil {
			return err
		}

		for _, performanceValue := range performanceValues {
			if performanceValue.PerformanceID == performance.ID {
				performance.PerformanceValue = append(performance.PerformanceValue, performanceValue)
			}
		}
		academic.PerformanceID = performance.ID
		academic.Performance = performance
		*academics = append(*academics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*academics))

	return nil
}
