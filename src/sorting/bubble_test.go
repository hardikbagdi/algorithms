package sorting

import "testing"
import "random"
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

func TestBubbleSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := random.RandomArray(10000)
		BubbleSort(arr)
		if !isIntArraySorted(arr) {
			t.Fatal("Failed")
		}
	}
}
