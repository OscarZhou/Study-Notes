package main

import "testing"

func TestFindData(t *testing.T) {
	err := FindData()
	if err != nil {
		t.Error(err)
	}

}
