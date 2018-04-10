package main

import (
	"fmt"
	"reflect"
	"strings"
)

type arrString []string

func ConvertMapToInSyntax(m interface{}) string {
	s := ""
	mk := reflect.ValueOf(m)
	if mk.IsValid() {
		switch mk.Kind() {
		case reflect.Map:
			mkeys := mk.MapKeys()
			fmt.Println(mkeys)
			for _, key := range mkeys {
				v := mk.MapIndex(key)
				s = key.String() + " IN "
				array, ok := v.Interface().([]string)
				if ok {
					var inSentence []string
					for _, v := range array {
						inSentence = append(inSentence, v)
					}
					s = s + "(" + strings.Join(inSentence, ",") + ")"
				}
			}
		}
	}
	return s
}

func main() {

}
