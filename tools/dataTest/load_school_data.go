package main

import (
	"Study-Notes/tools/dataTest/models"
	"fmt"
	"reflect"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func LoadSchoolData() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	db.DropTable(
		models.NZSchoolAcademic{},
	)

	db.AutoMigrate(
		models.NZSchoolAcademic{},
	)

	// db.LogMode(true)

	var schoolAcademics []models.NZSchoolAcademic
	// 1
	if err = loadSchoolAcademicEthnicityEndorsement(db, &schoolAcademics); err != nil {
		return err
	}

	// 2
	if err = loadSchoolAcademicEthnicityLit(db, &schoolAcademics); err != nil {
		return err
	}
	// 3
	if err = loadSchoolAcademicEthnicityQualification(db, &schoolAcademics); err != nil {
		return err
	}

	// 4
	if err = loadSchoolAcademicGenderEndorsement(db, &schoolAcademics); err != nil {
		return err
	}
	// 5
	if err = loadSchoolAcademicGenderLit(db, &schoolAcademics); err != nil {
		return err
	}

	// 6
	if err = loadSchoolAcademicGenderQualification(db, &schoolAcademics); err != nil {
		return err
	}

	err = batchCreate(db, schoolAcademics)
	if err != nil {
		return err
	}
	// 7
	if err = loadNCEASchoolAcademicEndorsement(db, &schoolAcademics, "Gender", "All"); err != nil {
		return err
	}
	// 8
	if err = loadNCEASchoolAcademicLit(db, &schoolAcademics, "Gender", "All"); err != nil {
		return err
	}

	// 9
	if err = loadNCEASchoolAcademicQualification(db, &schoolAcademics, "Gender", "All"); err != nil {
		return err
	}

	// 10
	if err = loadNCEASchoolAcademicEndorsement(db, &schoolAcademics, "Ethnicity", "All"); err != nil {
		return err
	}
	// 11
	if err = loadNCEASchoolAcademicLit(db, &schoolAcademics, "Ethnicity", "All"); err != nil {
		return err
	}

	// 12
	if err = loadNCEASchoolAcademicQualification(db, &schoolAcademics, "Ethnicity", "All"); err != nil {
		return err
	}
	// 14
	if err = loadNCEASchoolAcademicEndorsement(db, &schoolAcademics, "National", "National"); err != nil {
		return err
	}
	// 15
	if err = loadNCEASchoolAcademicLit(db, &schoolAcademics, "National", "National"); err != nil {
		return err
	}

	// 16
	if err = loadNCEASchoolAcademicQualification(db, &schoolAcademics, "National", "National"); err != nil {
		return err
	}

	err = batchCreate(db, schoolAcademics)
	if err != nil {
		return err
	}
	return nil
}

func batchCreate(db *gorm.DB, schoolAcademics []models.NZSchoolAcademic) error {
	insertSQL := `INSERT INTO nz_school_academics (school_id, year, year_level_id, category_id, performance_id, current_achievement_rate, cumulative_achievement_rate, decile_band) VALUES `
	for i, v := range schoolAcademics {
		insertSQL += ("(" + strconv.FormatUint(uint64(v.SchoolID), 10) + ", " + strconv.FormatUint(uint64(v.Year), 10) + "," + strconv.FormatUint(uint64(v.YearLevelID), 10) + "," + strconv.FormatUint(uint64(v.CategoryID), 10) + "," + strconv.FormatUint(uint64(v.PerformanceID), 10) + "," + strconv.FormatFloat(v.CurrentAchievementRate, 'f', 4, 64) + "," + strconv.FormatFloat(v.CumulativeAchievementRate, 'f', 4, 64) + ", '" + v.DecileBand + "'),")
		if (i+1)%1000 == 0 {
			err := db.Exec(insertSQL[0 : len(insertSQL)-1]).Error
			if err != nil {
				return err
			}
			insertSQL = `INSERT INTO nz_school_academics (school_id, year, year_level_id, category_id, performance_id, current_achievement_rate, cumulative_achievement_rate, decile_band) VALUES `
		}
		if i+1 == len(schoolAcademics) {
			fmt.Println("==", i)
			err := db.Exec(insertSQL[0 : len(insertSQL)-1]).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func loadSchoolAcademicEthnicityEndorsement(db *gorm.DB, schoolAcademics *[]models.NZSchoolAcademic) error {
	var (
		entities []models.SchoolAcademicEthnicityEndorsement
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var (
		yealLevels []models.NZYearLevel
		categories []models.NZCategory
	)
	err = db.Find(&yealLevels).Error
	if err != nil {
		return err
	}

	// DecileBand
	err = db.Where("name = ?", "Ethnicity").Find(&categories).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			academic    models.NZSchoolAcademic
			performance models.NZPerformance
		)

		academic.Year = v.Year
		academic.SchoolID = v.SchoolID
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		switch v.DecileBand {
		case "1", "2", "3":
			academic.DecileBand = "Decile 1-3"
		case "4", "5", "6", "7":
			academic.DecileBand = "Decile 4-7"
		case "8", "9", "10":
			academic.DecileBand = "Decile 8-10"
		case "0", "99":
			academic.DecileBand = "Unknown"
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		for _, category := range categories {
			if category.Value == v.Ethnicity {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}

		err = db.Raw(`select ip.* 
			from info_nz_performances ip 
			join info_nz_performance_values ipv1 on ipv1.performance_id = ip.id 
			join info_nz_performance_values ipv2 on ipv2.performance_id = ip.id 
			where ip.name = ? and ipv1.name = ? and ipv2.name = ?`, "Endorsement",
			v.Endorsement, v.Qualification).Scan(&performance).Error
		if err != nil {
			return err
		}

		academic.Performance = performance
		*schoolAcademics = append(*schoolAcademics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*schoolAcademics))

	return nil
}

func loadSchoolAcademicEthnicityLit(db *gorm.DB, schoolAcademics *[]models.NZSchoolAcademic) error {
	var (
		entities []models.SchoolAcademicEthnicityLit
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var (
		yealLevels []models.NZYearLevel
		categories []models.NZCategory
	)
	err = db.Find(&yealLevels).Error
	if err != nil {
		return err
	}

	// DecileBand
	err = db.Where("name = ?", "Ethnicity").Find(&categories).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			academic    models.NZSchoolAcademic
			performance models.NZPerformance
		)

		academic.Year = v.Year
		academic.SchoolID = v.SchoolID
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		switch v.DecileBand {
		case "1", "2", "3":
			academic.DecileBand = "Decile 1-3"
		case "4", "5", "6", "7":
			academic.DecileBand = "Decile 4-7"
		case "8", "9", "10":
			academic.DecileBand = "Decile 8-10"
		case "0", "99":
			academic.DecileBand = "Unknown"
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		for _, category := range categories {
			if category.Value == v.Ethnicity {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}
		err = db.Raw(`select ip.* 
			from info_nz_performances ip 
			join info_nz_performance_values ipv1 on ipv1.performance_id = ip.id 
			where ip.name = ? and ipv1.name = ? `, "Lit",
			v.Certificate).Scan(&performance).Error
		if err != nil {
			return err
		}
		academic.Performance = performance
		*schoolAcademics = append(*schoolAcademics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*schoolAcademics))

	return nil
}

func loadSchoolAcademicEthnicityQualification(db *gorm.DB, schoolAcademics *[]models.NZSchoolAcademic) error {
	var (
		entities []models.SchoolAcademicEthnicityQualification
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var (
		yealLevels []models.NZYearLevel
		categories []models.NZCategory
	)
	err = db.Find(&yealLevels).Error
	if err != nil {
		return err
	}

	// DecileBand
	err = db.Where("name = ?", "Ethnicity").Find(&categories).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			academic    models.NZSchoolAcademic
			performance models.NZPerformance
		)

		academic.Year = v.Year
		academic.SchoolID = v.SchoolID
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		switch v.DecileBand {
		case "1", "2", "3":
			academic.DecileBand = "Decile 1-3"
		case "4", "5", "6", "7":
			academic.DecileBand = "Decile 4-7"
		case "8", "9", "10":
			academic.DecileBand = "Decile 8-10"
		case "0", "99":
			academic.DecileBand = "Unknown"
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		for _, category := range categories {
			if category.Value == v.Ethnicity {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}
		err = db.Raw(`select ip.* 
			from info_nz_performances ip 
			join info_nz_performance_values ipv1 on ipv1.performance_id = ip.id 
			where ip.name = ? and ipv1.name = ? `, "Qualification",
			v.Qualification).Scan(&performance).Error
		if err != nil {
			return err
		}

		academic.Performance = performance
		*schoolAcademics = append(*schoolAcademics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*schoolAcademics))

	return nil
}

func loadSchoolAcademicGenderEndorsement(db *gorm.DB, schoolAcademics *[]models.NZSchoolAcademic) error {
	var (
		entities []models.SchoolAcademicGenderEndorsement
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var (
		yealLevels []models.NZYearLevel
		categories []models.NZCategory
	)
	err = db.Find(&yealLevels).Error
	if err != nil {
		return err
	}

	// DecileBand
	err = db.Where("name = ?", "Gender").Find(&categories).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			academic    models.NZSchoolAcademic
			performance models.NZPerformance
		)

		academic.Year = v.Year
		academic.SchoolID = v.SchoolID
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		switch v.DecileBand {
		case "1", "2", "3":
			academic.DecileBand = "Decile 1-3"
		case "4", "5", "6", "7":
			academic.DecileBand = "Decile 4-7"
		case "8", "9", "10":
			academic.DecileBand = "Decile 8-10"
		case "0", "99":
			academic.DecileBand = "Unknown"
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		for _, category := range categories {
			if category.Value == v.Gender {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}
		err = db.Raw(`select ip.* 
			from info_nz_performances ip 
			join info_nz_performance_values ipv1 on ipv1.performance_id = ip.id 
			join info_nz_performance_values ipv2 on ipv2.performance_id = ip.id 
			where ip.name = ? and ipv1.name = ? and ipv2.name = ?`, "Endorsement",
			v.Endorsement, v.Qualification).Scan(&performance).Error
		if err != nil {
			return err
		}
		academic.Performance = performance
		*schoolAcademics = append(*schoolAcademics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*schoolAcademics))

	return nil
}

func loadSchoolAcademicGenderLit(db *gorm.DB, schoolAcademics *[]models.NZSchoolAcademic) error {
	var (
		entities []models.SchoolAcademicGenderLit
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var (
		yealLevels []models.NZYearLevel
		categories []models.NZCategory
	)
	err = db.Find(&yealLevels).Error
	if err != nil {
		return err
	}

	// DecileBand
	err = db.Where("name = ?", "Gender").Find(&categories).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			academic    models.NZSchoolAcademic
			performance models.NZPerformance
		)

		academic.Year = v.Year
		academic.SchoolID = v.SchoolID
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		switch v.DecileBand {
		case "1", "2", "3":
			academic.DecileBand = "Decile 1-3"
		case "4", "5", "6", "7":
			academic.DecileBand = "Decile 4-7"
		case "8", "9", "10":
			academic.DecileBand = "Decile 8-10"
		case "0", "99":
			academic.DecileBand = "Unknown"
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		for _, category := range categories {
			if category.Value == v.Gender {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}
		err = db.Raw(`select ip.* 
			from info_nz_performances ip 
			join info_nz_performance_values ipv1 on ipv1.performance_id = ip.id 
			where ip.name = ? and ipv1.name = ? `, "Lit",
			v.Certificate).Scan(&performance).Error
		if err != nil {
			return err
		}
		academic.PerformanceID = performance.ID
		*schoolAcademics = append(*schoolAcademics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*schoolAcademics))

	return nil
}

func loadSchoolAcademicGenderQualification(db *gorm.DB, schoolAcademics *[]models.NZSchoolAcademic) error {
	var (
		entities []models.SchoolAcademicGenderQualification
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var (
		yealLevels []models.NZYearLevel
		categories []models.NZCategory
	)
	err = db.Find(&yealLevels).Error
	if err != nil {
		return err
	}

	// DecileBand
	err = db.Where("name = ?", "Gender").Find(&categories).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			academic    models.NZSchoolAcademic
			performance models.NZPerformance
		)

		academic.Year = v.Year
		academic.SchoolID = v.SchoolID
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate

		switch v.DecileBand {
		case "1", "2", "3":
			academic.DecileBand = "Decile 1-3"
		case "4", "5", "6", "7":
			academic.DecileBand = "Decile 4-7"
		case "8", "9", "10":
			academic.DecileBand = "Decile 8-10"
		case "0", "99":
			academic.DecileBand = "Unknown"
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		for _, category := range categories {
			if category.Value == v.Gender {
				academic.CategoryID = category.ID
				academic.Category = category
				break
			}
		}
		err = db.Raw(`select ip.* 
			from info_nz_performances ip 
			join info_nz_performance_values ipv1 on ipv1.performance_id = ip.id 
			where ip.name = ? and ipv1.name = ? `, "Qualification",
			v.Qualification).Scan(&performance).Error
		if err != nil {
			return err
		}
		academic.PerformanceID = performance.ID
		*schoolAcademics = append(*schoolAcademics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*schoolAcademics))

	return nil
}

func loadNCEASchoolAcademicEndorsement(db *gorm.DB, schoolAcademics *[]models.NZSchoolAcademic, cateName, cateValue string) error {
	var (
		entities []models.NCEASchoolAcademicEndorsement
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var (
		yealLevels []models.NZYearLevel
		category   models.NZCategory
	)
	err = db.Find(&yealLevels).Error
	if err != nil {
		return err
	}

	// DecileBand
	err = db.Where("name = ? and value = ?", cateName, cateValue).Find(&category).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			academic    models.NZSchoolAcademic
			performance models.NZPerformance
		)

		academic.Year = v.Year
		academic.SchoolID = v.SchoolID
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate
		academic.CategoryID = category.ID

		switch v.Decile {
		case "1", "2", "3":
			academic.DecileBand = "Decile 1-3"
		case "4", "5", "6", "7":
			academic.DecileBand = "Decile 4-7"
		case "8", "9", "10":
			academic.DecileBand = "Decile 8-10"
		case "0", "99":
			academic.DecileBand = "Unknown"
		}
		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}

		err = db.Raw(`select ip.* 
			from info_nz_performances ip 
			join info_nz_performance_values ipv1 on ipv1.performance_id = ip.id 
			join info_nz_performance_values ipv2 on ipv2.performance_id = ip.id 
			where ip.name = ? and ipv1.name = ? and ipv2.name = ?`, "Endorsement",
			v.Endorsement, v.Qualification).Scan(&performance).Error
		if err != nil {
			return err
		}

		academic.PerformanceID = performance.ID
		*schoolAcademics = append(*schoolAcademics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*schoolAcademics))

	return nil
}

func loadNCEASchoolAcademicLit(db *gorm.DB, schoolAcademics *[]models.NZSchoolAcademic, cateName, cateValue string) error {
	var (
		entities []models.NCEASchoolAcademicLit
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var (
		yealLevels []models.NZYearLevel
		category   models.NZCategory
	)
	err = db.Find(&yealLevels).Error
	if err != nil {
		return err
	}

	// DecileBand
	err = db.Where("name = ? and value = ?", cateName, cateValue).Find(&category).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			academic    models.NZSchoolAcademic
			performance models.NZPerformance
		)

		academic.Year = v.Year
		academic.SchoolID = v.SchoolID
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate
		academic.CategoryID = category.ID

		switch v.Decile {
		case "1", "2", "3":
			academic.DecileBand = "Decile 1-3"
		case "4", "5", "6", "7":
			academic.DecileBand = "Decile 4-7"
		case "8", "9", "10":
			academic.DecileBand = "Decile 8-10"
		case "0", "99":
			academic.DecileBand = "Unknown"
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}
		err = db.Raw(`select ip.* 
			from info_nz_performances ip 
			join info_nz_performance_values ipv1 on ipv1.performance_id = ip.id 
			where ip.name = ? and ipv1.name = ? `, "Lit",
			v.Certificate).Scan(&performance).Error
		if err != nil {
			return err
		}

		academic.PerformanceID = performance.ID
		*schoolAcademics = append(*schoolAcademics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*schoolAcademics))

	return nil
}

func loadNCEASchoolAcademicQualification(db *gorm.DB, schoolAcademics *[]models.NZSchoolAcademic, cateName, cateValue string) error {
	var (
		entities []models.NCEASchoolAcademicQualification
		err      error
	)

	if err = db.Find(&entities).Error; err != nil {
		return err
	}

	var (
		yealLevels []models.NZYearLevel
		category   models.NZCategory
	)
	err = db.Find(&yealLevels).Error
	if err != nil {
		return err
	}

	// DecileBand
	err = db.Where("name = ? and value = ?", cateName, cateValue).Find(&category).Error
	if err != nil {
		return err
	}

	for _, v := range entities {
		var (
			academic    models.NZSchoolAcademic
			performance models.NZPerformance
		)

		academic.Year = v.Year
		academic.SchoolID = v.SchoolID
		academic.CurrentAchievementRate = v.CurrentAchievementRate
		academic.CumulativeAchievementRate = v.CumulativeAchievementRate
		academic.CategoryID = category.ID

		switch v.Decile {
		case "1", "2", "3":
			academic.DecileBand = "Decile 1-3"
		case "4", "5", "6", "7":
			academic.DecileBand = "Decile 4-7"
		case "8", "9", "10":
			academic.DecileBand = "Decile 8-10"
		case "0", "99":
			academic.DecileBand = "Unknown"
		}

		for _, yearLevel := range yealLevels {
			if yearLevel.Level == v.YearLevel {
				academic.YearLevelID = yearLevel.ID
				academic.YearLevel = yearLevel
				break
			}
		}
		err = db.Raw(`select ip.* 
			from info_nz_performances ip 
			join info_nz_performance_values ipv1 on ipv1.performance_id = ip.id 
			where ip.name = ? and ipv1.name = ? `, "Qualification",
			v.Qualification).Scan(&performance).Error
		if err != nil {
			return err
		}
		academic.PerformanceID = performance.ID
		*schoolAcademics = append(*schoolAcademics, academic)
	}

	fmt.Printf("table:%s, %d, %d\n", reflect.ValueOf(entities[0]).Type().Name(), len(entities), len(*schoolAcademics))

	return nil
}
