package main

import "testing"

func TestLoadSchoolData(t *testing.T) {
	if err := LoadSchoolData(); err != nil {
		panic(err)
	}
}
