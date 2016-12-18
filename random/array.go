// Package random is a wrapper around golang's math/rand
package random

import "math/rand"
import "time"

// RandomizeSeed sets a random seed
func RandomizeSeed() {
	t := time.Now()
	rand.Seed(t.Unix())
}

// Array allocate a new random array of size 'size with elments 0-999
func Array(size int) []int {
	RandomizeSeed()
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	return arr
}
