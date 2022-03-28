package main

import (
	"fmt"
	"sort"
)

// Given five positive integers, find the minimum and maximum values that can be calculated
// by summing exactly four of the five integers. Then print the respective minimum and maximum
// values as a single line of two space-separated long integers.

func minMaxSum(arr []int32) {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var minSum, maxSum int

	for i := 1; i < len(arr); i++ {
		maxSum += int(arr[i])
	}
	for i := 0; i < len(arr)-1; i++ {
		minSum += int(arr[i])
	}

	fmt.Printf("%d %d", minSum, maxSum)
}

func main() {
	testSlice := []int32{1, 2, 3, 4, 5}
	minMaxSum(testSlice)
}
