package main

import (
	"fmt"
)

func main() {
	s := `{
		"size":"1",
		"query":{
		   "filtered":{
			  "filter":{
				 "bool":{
					"must":[
					   {
						  "term":{
							 "Status":"CurrencyActivated"
						  }
					   },
					   {
						"term":{
						   "Code":"USD"
						}
					 }
					]
				 }
			  }
		   }
		}
	 }
	 `
	// ss := strings.Replace(s, "", "", -1)

	// fmt.Println(strings.Replace(ss, " ", "", -1))

	fmt.Println(removeSpaceAndNextLine(s))
}

func removeSpaceAndNextLine(s string) string {
	var newString []byte
	for i := 0; i < len(s); i++ {

		if (s[i] == '\n') || (s[i] == ' ') || (s[i] == '\t') {
			continue
		}
		newString = append(newString, byte(s[i]))
	}
	return string(newString)
}
