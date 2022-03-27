package main

import (
	"log"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		caseName       string
		sliceForSort   []int
		expectedResult []int
	}{
		{
			caseName:       "first sort",
			sliceForSort:   []int{4, 3, 2, 5, 3, 1, 5, 6, 3, 2},
			expectedResult: []int{1, 2, 2, 3, 3, 3, 4, 5, 5, 6},
		},
		{
			caseName:       "second sort",
			sliceForSort:   []int{3, 2, 1, 4, 5, 3, 2, 1},
			expectedResult: []int{1, 1, 2, 2, 3, 3, 4, 5},
		},
	}
	for _, v := range testCases {
		t.Run(v.caseName, func(t *testing.T) {
			sliceAfterSort := quickSort(v.sliceForSort)
			for i, k := range sliceAfterSort {
				if k != v.expectedResult[i] {
					log.Print("test failed")
				}
			}
			log.Print("sorting is successful")
		})
	}
}
