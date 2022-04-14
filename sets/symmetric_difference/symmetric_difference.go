package main

import "fmt"

func SymmetricDifference(firstArr, secondArr []int) []int {
	var result []int

	for _, v := range firstArr {
		if !includedInArr(v, secondArr) {
			result = append(result, v)
		}
	}

	for _, v := range secondArr {
		if !includedInArr(v, firstArr) {
			result = append(result, v)
		}
	}

	return result
}

func includedInArr(num int, arr []int) bool {
	for _, v := range arr {
		if num == v {
			return true
		}
	}
	return false
}

func main() {
	firstArr := []int{1, 2, 3, 4, 5}
	secondArr := []int{4, 5, 6, 7, 8}

	fmt.Println(SymmetricDifference(firstArr, secondArr)) // symmetric difference is [1, 2, 3, 6, 7, 8]
}
