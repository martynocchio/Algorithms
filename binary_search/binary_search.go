package main

import (
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
