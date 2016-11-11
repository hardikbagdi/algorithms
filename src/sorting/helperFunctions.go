package sorting

import "sort"
import "fmt"
import "math"

func arraysEqual(arr1, arr2 []int) bool {

	for i, val := range arr1 {
		if val != arr2[i] {
			return false
		}
	}
	return true
}

func isIntArraySorted(arr []int) bool {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sort.Ints(arrCopy)
	return arraysEqual(arr, arrCopy)
}

func xchg(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func printArray(array []int) {
	fmt.Println("Array: ", array, " length: ", len(array))
}

func getMax(array []int) (max int) {

	max = math.MinInt64
	for _, val := range array {
		if max < val {
			max = val
		}
	}
	return
}
