package main

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	lessSlice := make([]int, 0)
	moreSlice := make([]int, 0)
	pivot := arr[0]

	for _, v := range arr[1:] {
		if v < pivot {
			lessSlice = append(lessSlice, v)
		} else {
			moreSlice = append(moreSlice, v)
		}
	}
	arr = append(quickSort(lessSlice), pivot)
	arr = append(arr, quickSort(moreSlice)...)
	return arr
}
