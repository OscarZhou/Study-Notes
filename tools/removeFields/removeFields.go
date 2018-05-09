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

	var u2 User
	u2.ID = "a002"
	u2.CreatedAt = time.Now()
	u2.UpdatedAt = time.Now()
	u2.Name = "creat"
	u2.Age = 23

	var us []User
	us = append(us, u)
	us = append(us, u2)

	fmt.Println("1: ", us)
	removeTimeFields(&us)
	fmt.Println("2: ", us)
}

func removeTimeFields(u interface{}) {
	values := reflect.ValueOf(u).Elem()
	ti := time.Time{}

	if values.Kind() == reflect.Slice {
		// fmt.Println("----------SLICE-----", values.Len())
		for k := 0; k < values.Len(); k++ {
			kv := values.Slice(k, k+1).Index(0)
			// fmt.Println("kv=", kv)
			for i := 0; i < kv.NumField(); i++ {
				t := kv.Type().Field(i).Type
				// fmt.Println(values.Type().Field(i).Type.Name())
				if t.Kind() == reflect.Struct && t.Name() == "Model" {
					vv := kv.Field(i)
					for j := 0; j < vv.NumField(); j++ {
						tt := vv.Type().Field(j).Type
						if tt.String() == "time.Time" {
							// fmt.Println("-----detect the time.Time type")
							k := kv.Field(i).Field(j)
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
	}

}
