package main

import (
	"Study-Notes/tools/questionaire_data/models"
	"dataant/dataant_go_api/api/models/types"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	if err := ImportQuestionaire(); err != nil {
		panic(err)
	}
}

func ImportQuestionaire() error {
	db, err := gorm.Open("postgres", "")
	defer db.Close()
	if err != nil {
		return err
	}

	db.DropTable(
		models.DynamicForm{},
		models.DynamicFormGroup{},
		models.DynamicFormGroupType{},
		models.DynamicFormQuestion{},
		models.DynamicFormOption{},
		models.AnalysisInput{},
	)

	db.AutoMigrate(
		models.DynamicForm{},
		models.DynamicFormGroup{},
		models.DynamicFormGroupType{},
		models.DynamicFormQuestion{},
		models.DynamicFormOption{},
		models.AnalysisInput{},
	)

	db.LogMode(true)

	err = addDynamicForm(db)
	if err != nil {
		return err
	}

	err = addDynamicFormGroupType(db)
	if err != nil {
		return err
	}

	err = addDynamicFormGroup(db)
	if err != nil {
		return err
	}

	err = addDynamicFormQuestion(db)
	if err != nil {
		return err
	}

	err = addDynamicFormOption(db)
	if err != nil {
		return err
	}

	return err
}

func addDynamicForm(db *gorm.DB) error {
	var dynamicForms []models.DynamicForm
	dynamicForms = append(dynamicForms, models.DynamicForm{ID: 1, Name: "Questionaire Survey", CountryCode: "NZ"})
	dynamicForms = append(dynamicForms, models.DynamicForm{ID: 2, Name: "Agency Application Form", CountryCode: "NZ"})
	dynamicForms = append(dynamicForms, models.DynamicForm{ID: 3, Name: "Questionaire Survey", CountryCode: "UK"})
	dynamicForms = append(dynamicForms, models.DynamicForm{ID: 4, Name: "Agency Application Form", CountryCode: "UK"})

	for _, v := range dynamicForms {
		err := db.Create(&v).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func addDynamicFormGroupType(db *gorm.DB) error {
	var groupTypes []models.DynamicFormGroupType
	groupTypes = append(groupTypes, models.DynamicFormGroupType{ID: 1, Name: "School Analysis"})
	groupTypes = append(groupTypes, models.DynamicFormGroupType{ID: 2, Name: "House Analysis"})
	for _, v := range groupTypes {
		err := db.Create(&v).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func addDynamicFormGroup(db *gorm.DB) error {
	var groups []models.DynamicFormGroup
	groups = append(groups, models.DynamicFormGroup{ID: 1, Name: "School Analysis", GroupTypeID: 1, DynamicFormID: 1})
	for _, v := range groups {
		err := db.Create(&v).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func addDynamicFormQuestion(db *gorm.DB) error {
	var questions []models.DynamicFormQuestion
	questions = append(questions, models.DynamicFormQuestion{ID: 1, GroupID: 1, Ordering: 1, Name: "Please choose the type of student",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutColumn4, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 2, GroupID: 1, Ordering: 2, Name: "What is the gender of your child?",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutColumn4, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 3, GroupID: 1, Ordering: 3, Name: "What is the age of your child?",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutColumn2, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 4, GroupID: 1, Ordering: 4, Name: "What is your preference of the school gender composition?",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutColumn4, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 5, GroupID: 1, Ordering: 5, Name: "Do you prefer to have your child study at the same school from primary to secondary?",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutColumn4, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 6, GroupID: 1, Ordering: 6, Name: "What type of school do you prefer?",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutColumn2, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 7, GroupID: 1, Ordering: 7, Name: "How important are the following foci to your child's learning academic performance",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutRuler, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 8, GroupID: 1, Ordering: 8, Name: "How important are the following foci to your child's learning instrument and voice",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutRuler, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 9, GroupID: 1, Ordering: 9, Name: "How important are the following foci to your child's learning sports",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutRuler, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 10, GroupID: 1, Ordering: 10, Name: "How important are the following foci to your child's learning art & performance",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutRuler, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 11, GroupID: 1, Ordering: 11, Name: "How important are the following foci to your child's learning language",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutRuler, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 12, GroupID: 1, Ordering: 12, Name: "What is the perspective of your child's higher education?",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutColumn4, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 13, GroupID: 1, Ordering: 13, Name: "School Region",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutColumn4, Online: true})
	questions = append(questions, models.DynamicFormQuestion{ID: 14, GroupID: 1, Ordering: 14, Name: "School Size",
		QuestionType: types.QuestionRadio, AnswerLayoutType: types.LayoutColumn4, Online: true})

	for _, v := range questions {
		err := db.Create(&v).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func addDynamicFormOption(db *gorm.DB) error {
	var options []models.DynamicFormOption
	options = append(options, models.DynamicFormOption{ID: 1, Ordering: 1, QuestionID: 1, Name: "International", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 2, Ordering: 2, QuestionID: 1, Name: "Domestic", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 3, Ordering: 1, QuestionID: 2, Name: "Girl", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 4, Ordering: 2, QuestionID: 2, Name: "Boy", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 5, Ordering: 1, QuestionID: 3, Name: "Primary Year 0-6 | Age 5-11", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 6, Ordering: 2, QuestionID: 3, Name: "Intermediate Year 7-8 | Age 12-13", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 7, Ordering: 3, QuestionID: 3, Name: "Secondary Year 9-10 | Age 14-15", Value: "3"})
	options = append(options, models.DynamicFormOption{ID: 8, Ordering: 4, QuestionID: 3, Name: "Secondary Year 11-13 | Age 16-18", Value: "4"})
	options = append(options, models.DynamicFormOption{ID: 9, Ordering: 1, QuestionID: 4, Name: "Single Gender", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 10, Ordering: 2, QuestionID: 4, Name: "Co-Edu", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 11, Ordering: 3, QuestionID: 4, Name: "Does not matter", Value: "3"})
	options = append(options, models.DynamicFormOption{ID: 12, Ordering: 1, QuestionID: 5, Name: "Does not matter", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 13, Ordering: 2, QuestionID: 5, Name: "Yes", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 14, Ordering: 1, QuestionID: 6, Name: "<7500(Most of Stated Schools)", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 15, Ordering: 2, QuestionID: 6, Name: "7500-15000(Most of Stated Integrate Schools)", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 16, Ordering: 3, QuestionID: 6, Name: "15000+(Most of Private Schools)", Value: "3"})
	options = append(options, models.DynamicFormOption{ID: 17, Ordering: 4, QuestionID: 6, Name: "Does not matter", Value: "4"})
	options = append(options, models.DynamicFormOption{ID: 18, Ordering: 1, QuestionID: 7, Name: "1", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 19, Ordering: 2, QuestionID: 7, Name: "2", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 20, Ordering: 3, QuestionID: 7, Name: "3", Value: "3"})
	options = append(options, models.DynamicFormOption{ID: 21, Ordering: 4, QuestionID: 7, Name: "4", Value: "4"})
	options = append(options, models.DynamicFormOption{ID: 22, Ordering: 5, QuestionID: 7, Name: "5", Value: "5"})
	options = append(options, models.DynamicFormOption{ID: 23, Ordering: 1, QuestionID: 8, Name: "1", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 24, Ordering: 2, QuestionID: 8, Name: "2", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 25, Ordering: 3, QuestionID: 8, Name: "3", Value: "3"})
	options = append(options, models.DynamicFormOption{ID: 26, Ordering: 4, QuestionID: 8, Name: "4", Value: "4"})
	options = append(options, models.DynamicFormOption{ID: 27, Ordering: 5, QuestionID: 8, Name: "5", Value: "5"})
	options = append(options, models.DynamicFormOption{ID: 28, Ordering: 1, QuestionID: 9, Name: "1", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 29, Ordering: 2, QuestionID: 9, Name: "2", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 30, Ordering: 3, QuestionID: 9, Name: "3", Value: "3"})
	options = append(options, models.DynamicFormOption{ID: 31, Ordering: 4, QuestionID: 9, Name: "4", Value: "4"})
	options = append(options, models.DynamicFormOption{ID: 32, Ordering: 5, QuestionID: 9, Name: "5", Value: "5"})
	options = append(options, models.DynamicFormOption{ID: 33, Ordering: 1, QuestionID: 10, Name: "1", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 34, Ordering: 2, QuestionID: 10, Name: "2", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 35, Ordering: 3, QuestionID: 10, Name: "3", Value: "3"})
	options = append(options, models.DynamicFormOption{ID: 36, Ordering: 4, QuestionID: 10, Name: "4", Value: "4"})
	options = append(options, models.DynamicFormOption{ID: 37, Ordering: 5, QuestionID: 10, Name: "5", Value: "5"})
	options = append(options, models.DynamicFormOption{ID: 38, Ordering: 1, QuestionID: 11, Name: "1", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 39, Ordering: 2, QuestionID: 11, Name: "2", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 40, Ordering: 3, QuestionID: 11, Name: "3", Value: "3"})
	options = append(options, models.DynamicFormOption{ID: 41, Ordering: 4, QuestionID: 11, Name: "4", Value: "4"})
	options = append(options, models.DynamicFormOption{ID: 42, Ordering: 5, QuestionID: 11, Name: "5", Value: "5"})
	options = append(options, models.DynamicFormOption{ID: 43, Ordering: 1, QuestionID: 12, Name: "NZ", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 44, Ordering: 2, QuestionID: 12, Name: "Overseas", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 45, Ordering: 1, QuestionID: 13, Name: "All NZ", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 46, Ordering: 2, QuestionID: 13, Name: "Only Auckland", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 47, Ordering: 3, QuestionID: 13, Name: "Only outside Auckland", Value: "3"})
	options = append(options, models.DynamicFormOption{ID: 48, Ordering: 1, QuestionID: 14, Name: "All", Value: "1"})
	options = append(options, models.DynamicFormOption{ID: 49, Ordering: 2, QuestionID: 14, Name: "Small(<100)", Value: "2"})
	options = append(options, models.DynamicFormOption{ID: 50, Ordering: 3, QuestionID: 14, Name: "Medium(100-450)", Value: "3"})
	options = append(options, models.DynamicFormOption{ID: 51, Ordering: 4, QuestionID: 14, Name: "Large(>450)", Value: "4"})
	for _, v := range options {
		err := db.Create(&v).Error
		if err != nil {
			return err
		}
	}
	return nil
}
