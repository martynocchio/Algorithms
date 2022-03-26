package main

import (
	"log"
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		arrForSort        []int
		expectedSortedArr []int
	}{
		{
			name:              "first test",
			arrForSort:        []int{4, 3, 2, 1, 5, 4, 2},
			expectedSortedArr: []int{1, 2, 2, 3, 4, 4, 5},
		},
		{
			name:              "second test",
			arrForSort:        []int{0, 3274623, 584, 23, 556, 2, 985, 63},
			expectedSortedArr: []int{0, 2, 23, 63, 556, 584, 985, 3274623},
		},
	}
	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			resultArr := mergeSort(v.arrForSort)
			for i, value := range resultArr {
				if value != v.expectedSortedArr[i] {
					log.Println("failed test")
					return
				}
			}
			log.Println("sorting is successful")
			log.Println(resultArr)
		})
	}
}
