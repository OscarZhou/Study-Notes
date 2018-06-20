package main

import "testing"

func TestSaveIntegratedData(t *testing.T) {
	err := SaveIntegratedData()
	if err != nil {
		t.Error(err)
	}
}
