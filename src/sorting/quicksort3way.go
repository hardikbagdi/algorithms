package sorting

import "math/rand"

//Sort an array of int using 3-way quick sort
func QuickSort3way(array []int) {
	if array == nil || len(array) <= 1 {
		return
	}
	lo := 0
	high := len(array) - 1
	pivotIndex := rand.Intn(len(array))
	Xchg(array, lo, pivotIndex)
	pivot := array[lo]
	i := lo
	gt := high
	lt := lo
	for i <= gt {
		if array[i] == pivot {
			i++
		} else if array[i] > pivot {
			Xchg(array, i, gt)
			gt--
		} else if array[i] < pivot {
			Xchg(array, i, lt)
			i++
			lt++
		}
	}

	QuickSort3way(array[:lt])
	QuickSort3way(array[gt+1:])
}
