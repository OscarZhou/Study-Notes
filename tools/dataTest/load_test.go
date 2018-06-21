package main

import "testing"

// func TestFindData(t *testing.T) {
// 	err := FindData()
// 	if err != nil {
// 		t.Error(err)
// 	}

// }
// func TestAcademicData(t *testing.T) {
// 	err := SaveIntegratedData()
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestSchoolAcademicData(t *testing.T) {
// 	err := LoadSchoolData()
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestLoadEntrancePercentage(t *testing.T) {
// 	if err := LoadEntrancePercentageData(); err != nil {
// 		panic(err)
// 	}
// }

// func Test2LevelCategory(t *testing.T) {
// 	if err := ImportCategory(); err != nil {
// 		panic(err)
// 	}
// }

func TestFillScholarshipYear(t *testing.T) {
	if err := FillScholarshipYear(); err != nil {
		panic(err)
	}
}
