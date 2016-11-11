package sorting

import "math/rand"

//sort and array of int using simple quick sort
func QuickSort(array []int) {
	if array == nil || len(array) <= 1 {
		return
	}
	length := len(array)
	pivotIndex := rand.Intn(length)
	Xchg(array, 0, pivotIndex)
	pivot := array[0]
	last := 0
	for i := 1; i < length; i++ {
		if array[i] <= pivot {
			last++
			Xchg(array, last, i)
		}
	}
	Xchg(array, 0, last)
	QuickSort(array[:last])
	QuickSort(array[last:])
}
