package main

import "fmt"

func intersection(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, v := range a {
		if _, ok := counter[v]; !ok {
			counter[v] = 1
		} else {
			counter[v] += 1
		}
	}

	for _, v := range b {
		if elem, ok := counter[v]; ok && elem > 0 {
			counter[v]--
			result = append(result, v)
		}
	}
	return result
}

func main() {
	a := []int{4, 3, 2, 1, 5}
	b := []int{4, 2, 1, 5}
	fmt.Println(intersection(a, b))
}
