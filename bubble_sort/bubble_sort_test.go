package main

import (
	"log"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("first sort", func(t *testing.T) {

		unsortedArray := []int{5, 3, 4, 2, 5, 1, 6}
		sortedArray := bubbleSort(unsortedArray)
		expectedAfterSortArray := []int{1, 2, 3, 4, 5, 5, 6}

		deepEqual := reflect.DeepEqual(sortedArray, expectedAfterSortArray)

		if deepEqual == false {
			log.Println("sorting failed")
			return
		}
		log.Println("sorting is successful")
		log.Println(sortedArray)
	})

	t.Run("second sort", func(t *testing.T) {

		unsortedArray := []int{4, 2, 1, 2, 3, 3, 4, 5}
		sortedArray := bubbleSort(unsortedArray)
		expectedAfterSortArray := []int{1, 2, 2, 3, 3, 4, 4, 5}

		for i, value := range sortedArray {
			if value != expectedAfterSortArray[i] {
				log.Println("sorting failed")
				return
			}
		}
		log.Println("sorting is successful")
		log.Println(sortedArray)
	})
}
