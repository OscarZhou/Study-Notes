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

// func TestFillScholarshipYear(t *testing.T) {
// 	if err := FillScholarshipYear(); err != nil {
// 		panic(err)
// 	}
// }

// func TestSchoolRollCategory(t *testing.T) {
// 	if err := ImportSchoolRollCategory(); err != nil {
// 		panic(err)
// 	}
// }

// func TestFilterSchoolZone(t *testing.T) {
// 	if err := FilterSchoolZone(); err != nil {
// 		panic(err)
// 	}
// }

// func TestFilterMeshblock(t *testing.T) {
// 	if err := FilterMeshblock(); err != nil {
// 		panic(err)
// 	}
// }
// func TestAddExtraCurriculumCategory(t *testing.T) {
// 	if err := AddExtraCurriculumCategory(); err != nil {
// 		panic(err)
// 	}
// }
func TestFillCategoryIDForExtraCurriculum(t *testing.T) {
	if err := FillCategoryID2ExtraCurriculum(); err != nil {
		panic(err)
	}
}
