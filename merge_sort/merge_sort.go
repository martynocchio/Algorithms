package main

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	return merge(mergeSort(arr[:len(arr)/2]), mergeSort(arr[len(arr)/2:]))
}

func merge(arrA, arrB []int) []int {
	i := 0
	j := 0
	result := make([]int, 0)

	for i < len(arrA) && j < len(arrB) {
		if arrA[i] < arrB[j] {
			result = append(result, arrA[i])
			i++
		} else {
			result = append(result, arrB[j])
			j++
		}
	}

	for ; i < len(arrA); i++ {
		result = append(result, arrA[i])
	}
	for ; j < len(arrB); j++ {
		result = append(result, arrB[j])
	}
	return result
}
