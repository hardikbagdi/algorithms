package random

import "math/rand"
import "time"

// Set a random seed
func RandomizeSeed() {
	t := time.Now()
	rand.Seed(t.Unix())
}

//allocate a new random array of size 'size with number 0-999
func RandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	return arr
}
