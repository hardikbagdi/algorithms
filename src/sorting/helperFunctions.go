package sorting

import "sort"
import "fmt"
import "math"

func ArraysEqual(arr1, arr2 []int) bool {

	for i, val := range arr1 {
		if val != arr2[i] {
			return false
		}
	}
	return true
}

func IsIntArraySorted(arr []int) bool {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sort.Ints(arrCopy)
	return ArraysEqual(arr, arrCopy)
}

func Xchg(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func PrintArray(array []int) {
	fmt.Println("Array: ", array, " length: ", len(array))
}

func GetMax(array []int) (max int) {

	max = math.MinInt64
	for _, val := range array {
		if max < val {
			max = val
		}
	}
	return
}
