package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	t.Run("first binary search", func(t *testing.T) {
		shouldExists := true
		iterationNumber := 3

		testSlice := []int{5, 3, 2, 6, 5, 3, 6, 8, 7, 6, 4, 7, 3, 6, 8, 9, 4, 6, 3}
		FindElement := 8
		exists, number := binarySearch(testSlice, FindElement)

		if exists != shouldExists || number != iterationNumber {
			fmt.Println("test failed")
		}
		fmt.Println("test completed successfully")
	})
}
