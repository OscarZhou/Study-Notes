package main

import (
	"fmt"
	"regexp"
)

func main() {
	testString := "SELECT * FROM users WHERE deleted_at IS NULL AND ((username= 'pig@zetaapp.co.nz')) ORDER BY  created_at DESC Limit 25"
	matched, err := regexp.MatchString("email|username|mobile", testString)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(matched)
}
