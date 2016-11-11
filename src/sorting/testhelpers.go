package sorting

import "sort"

func arraysEqual(arr1, arr2 []int) bool {

	for i, val := range arr1 {
		if val != arr2[i] {
			return false
		}
	}
	return true
}

func isIntArraySorted(arr []int) bool {
	arrCopy := arr
	sort.Ints(arrCopy)
	return arraysEqual(arr, arrCopy)
}
