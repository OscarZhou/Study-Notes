package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapToInSyntax(t *testing.T) {
	m := make(map[string][]string)
	testData := `
	{
		"CurrencyCode":["LKR","NZD","USD"]
	}	
	`

	test := "CurrencyCode IN (LKR,NZD,USD)"
	err := json.Unmarshal([]byte(testData), &m)
	if err != nil {
		fmt.Println(err)
	}
	result := ConvertMapToInSyntax(m)
	if result != test {
		t.Errorf("get %s, want %s", result, test)
	}
}
