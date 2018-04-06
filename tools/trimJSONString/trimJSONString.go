package main

import (
	"fmt"
)

func TrimJSONString(s string) string {
	var newString []byte
	for i := 0; i < len(s); i++ {

		if (s[i] == '\n') || (s[i] == ' ') || (s[i] == '\t') {
			continue
		}
		newString = append(newString, byte(s[i]))
	}
	return string(newString)
}

func main() {
	testDataBefore := `{
		"query":{
		   "filtered":{
			  "filter":{
				 "bool":{
					"must":[
					   {
						  "term":{
							 "Status":"CurrencyActivated"
						  }
					   }
					]
				 }
			  }
		   }
		},
		"sort":[
		   {
			  "order":{
				 "Code":"desc",
				 "CreatedAt":"desc"
			  }
		   }
		]
	 }`

	result := TrimJSONString(testDataBefore)
	fmt.Println(result)
}
