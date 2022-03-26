package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, elem int) (exists bool, num int) {
	sort.Ints(arr)

	firstIndex := 0
	lastIndex := len(arr) - 1
	num = 0

	for firstIndex <= lastIndex {
		middleIndex := (firstIndex + lastIndex) / 2
		num++
		if arr[middleIndex] == elem {
			return true, num
		} else if arr[middleIndex] > elem {
			lastIndex = middleIndex - 1
		} else if arr[middleIndex] < elem {
			firstIndex = middleIndex + 1
		}
	}
	return false, 0
}

func main() {
	testArray := []int{5, 3, 2, 6, 5, 3, 6, 8, 7, 6, 4, 7, 3, 6, 8, 9, 4, 6, 3}
	fmt.Println(binarySearch(testArray, 8))
}
