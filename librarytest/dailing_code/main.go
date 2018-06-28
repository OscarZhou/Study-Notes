package main

import (
	"fmt"
	"regexp"

	"github.com/davegardnerisme/phonegeocode"
	"github.com/dongri/phonenumber"
)

func main() {
	mmobile := "+640211348864"
	cc, err := phonegeocode.New().Country(mmobile)

	if err != nil {
		panic(err)
	}

	fmt.Println(cc)
	mobile := ""
	countryCode := ""
	for _, v := range phonenumber.GetISO3166() {
		if v.Alpha2 == cc {
			mobile = regexp.MustCompile(`\D`+v.CountryCode).ReplaceAllString(mmobile, "")
			countryCode = "+" + v.CountryCode
			break
		}
	}

	p := "+" + phonenumber.Parse(mobile, cc)

	fmt.Println(p)

	fmt.Println(countryCode)

}
