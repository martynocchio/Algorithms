package main

import (
	"fmt"
)

//  Given an array of integers, calculate the ratios of its elements that are positive,
//	negative, and zero. Print the decimal value of each fraction
//	on a new line with 6 places after the decimal.

func calculateRatio(arr []int) {
	pos := 0
	neg := 0
	null := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			pos++
		} else if arr[i] == 0 {
			null++
		} else {
			neg++
		}
	}

	arrLen := len(arr)
	posRatio := float64(pos) / float64(arrLen)
	negRatio := float64(neg) / float64(arrLen)
	nullRatio := float64(null) / float64(arrLen)

	fmt.Printf("%.6f\n", posRatio)
	fmt.Printf("%.6f\n", negRatio)
	fmt.Printf("%.6f\n", nullRatio)
}

func main() {
	testSlice := []int{1, 2, -2, -6, 0, 2}
	calculateRatio(testSlice)
}
