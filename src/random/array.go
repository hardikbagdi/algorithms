package random

import "math/rand"

//allocate a new random array of size 'size with number 0-999
func RandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	return arr
}
