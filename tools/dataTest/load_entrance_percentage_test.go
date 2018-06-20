package main

import "testing"

func TestLoadEntrancePercentage(t *testing.T) {
	if err := LoadEntrancePercentageData(); err != nil {
		panic(err)
	}
}
