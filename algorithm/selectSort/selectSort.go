package main

import (
	"fmt"
)

func main() {
	array := []int{1, 34, 3, 69, 5, 55, 7, 21, 9, 91, 12, 59, 14, 15, 16}

	fmt.Println("old array: ", array)

	selectSort(&array)
	fmt.Println("new array: ", array)
}

func selectSort(arr *[]int) {
	var newArr []int
	arrCopy := *arr
	for i := 0; i < len(*arr); i++ {
		pos := findSmallest(arrCopy)
		newArr = append(newArr, arrCopy[pos])
		arrCopy = append(arrCopy[0:pos], arrCopy[pos+1:]...)
	}

	*arr = newArr
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}

	return smallestIndex
}
