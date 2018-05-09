package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	example := "#Go  Lang Code!$!"

	fmt.Printf("A string of %s becomes %s \n", example, GeneratePath(example))

	e := "Cars & Other Vehicles"
	fmt.Printf("A string of %s becomes %s \n", e, GeneratePath(e))
}

func GeneratePath(original string) string {
	// Make a Regex to say we only want
	reg, err := regexp.Compile(`[^a-zA-Z0-9\s]+`)
	if err != nil {
		log.Fatal(err)
	}
	p := strings.ToLower(reg.ReplaceAllString(original, ""))

	a := strings.Split(p, " ")

	var x []string
	for _, v := range a {
		if v != "" {
			x = append(x, v)
		}
	}

	return strings.Join(x, "-")

}
