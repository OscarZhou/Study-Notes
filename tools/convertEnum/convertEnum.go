package main

import (
	"FEI/fei_go_library/models/types"
	"encoding/json"
	"fmt"
)

type UserStatus int

const (
	Blocked UserStatus = iota
	Unblocked
)

type s struct {
	Status types.ArticleStatusType
	// Status UserStatus
}

func main() {
	var i s
	i.Status = types.ArticlePublished
	fmt.Println(i.Status)

	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println("marshal", err)
	}
	fmt.Println(string(b))

	// k := string(b)

	ii := &struct {
		Status types.ArticleStatusType
	}{}

	bb := []byte(`{"Status":"ArticlePublished"}`)
	err = json.Unmarshal(bb, ii)
	if err != nil {
		fmt.Println("unmarshal", err)
	}

	fmt.Println(ii)
}
