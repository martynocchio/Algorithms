package main

import "fmt"

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				arr[j], arr[min] = arr[min], arr[j]
			}
		}
	}
	return arr
}

func main() {
	testSlice := []int{6, 4, 2, 4, 5, 3, 2, 1, 5, 7, 8, 9, 6, 4, 3, 0}
	fmt.Println(selectionSort(testSlice))
}
