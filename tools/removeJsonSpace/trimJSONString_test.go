package main

import "testing"

func TestTrimJSONString(t *testing.T) {
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
				 "CreatedAt":"desc",
				 "Code":"desc"
			  }
		   }
		]
	 }`

	testDataAfter := `{"query":{"filtered":{"filter":{"bool":{"must":[{"term":{"Status":"CurrencyActivated"}}]}}}},"sort":[{"order":{"CreatedAt":"desc","Code":"desc"}}]}`

	result := TrimJSONString(testDataBefore)

	if result != testDataAfter {
		t.Errorf("test data's result get %s, want %s", result, testDataAfter)
	}

	t.Log("testing successfully!")
}
