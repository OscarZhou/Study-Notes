package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16}

	fmt.Println(BinarySearch(5, array))
}

func BinarySearch(target int, arr []int) (pos int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (right + left) / 2
		guess := arr[mid]
		if target == guess {
			return mid
		}
		if target < guess {
			right = mid - 1
			continue
		}
		if target > guess {
			left = mid + 1
			continue
		}
	}
	return -1
}
