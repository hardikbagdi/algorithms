package sorting

import (
	"math"
)

//Merge sort an array of int
func MergeSort(array []int) {
	//base case + some error checking
	if array == nil || len(array) <= 1 {
		return
	}
	mid := int(math.Floor(float64(len(array)) / 2))
	//recurseive calls
	MergeSort(array[:mid])
	MergeSort(array[mid:])
	//merge the two sorted arrays

	merge(array, 0, mid)
}

func merge(array []int, leftStartIndex, rightStartIndex int) {
	if array == nil {
		return
	}
	//make a copy of the two arrays to be merged
	lenLeft := rightStartIndex - leftStartIndex
	lenRight := len(array) - rightStartIndex
	left := make([]int, lenLeft)
	right := make([]int, lenRight)
	copy(left, array[:rightStartIndex])
	copy(right, array[rightStartIndex:])
	i, j, k := 0, 0, 0

	//merge
	for i < lenLeft && j < lenRight {
		if left[i] < right[j] {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}
		k++
	}
	//copy the remaining elements
	for i < lenLeft {
		array[k] = left[i]
		k++
		i++
	}
	for j < lenRight {
		array[k] = right[j]
		k++
		j++
	}
}
