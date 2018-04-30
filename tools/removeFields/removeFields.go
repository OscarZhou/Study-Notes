package main

import (
	"fmt"
	"reflect"
	"time"
)

type Model struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type User struct {
	Model
	Name string `json:"name"`
	Age  int64  `json:"-"`
}

func main() {
	var u User
	u.ID = "a001"
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.Name = "lala"
	u.Age = 21

	fmt.Println("1: ", u)
	removeTimeFields(&u)
	fmt.Println("2: ", u)
}

func removeTimeFields(u interface{}) {
	values := reflect.ValueOf(u).Elem()
	ti := time.Time{}
	for i := 0; i < values.NumField(); i++ {
		t := values.Type().Field(i).Type
		// fmt.Println(values.Type().Field(i).Type.Name())
		if t.Kind() == reflect.Struct && t.Name() == "Model" {
			vv := values.Field(i)
			for j := 0; j < vv.NumField(); j++ {
				tt := vv.Type().Field(j).Type
				if tt.String() == "time.Time" {
					// fmt.Println("-----detect the time.Time type")
					k := values.Field(i).Field(j)
					k.Set(reflect.ValueOf(ti))
					// fmt.Println(k)
					continue
				}
				// tv := vv.Field(j)
				// fmt.Println(tv, tt)
			}
			continue
		}
		// v := values.Field(i)
		// fmt.Println(v, t)
	}

}
