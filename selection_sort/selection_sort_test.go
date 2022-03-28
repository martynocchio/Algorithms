package main

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		sliceForSort   []int
		expectedResult []int
	}{
		{
			name:           "first sort",
			sliceForSort:   []int{4, 2, 1, 4, 6, 3, 2, 5, 3},
			expectedResult: []int{1, 2, 2, 3, 3, 4, 4, 5, 6},
		},
		{
			name:           "second sort",
			sliceForSort:   []int{6, 4, 2, 4, 5, 3, 2, 1, 5, 7, 8, 9, 6, 4, 3, 0},
			expectedResult: []int{0, 1, 2, 2, 3, 3, 4, 4, 4, 5, 5, 6, 6, 7, 8, 9},
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			sliceAfterSort := selectionSort(v.sliceForSort)
			for i, e := range sliceAfterSort {
				if e != v.expectedResult[i] {
					fmt.Println("test failed")
				}
			}
			fmt.Println("sorting is successful")
		})
	}
}
