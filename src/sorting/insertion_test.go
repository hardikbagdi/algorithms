package sorting

import (
	"random"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 10; i++ {
		arr := random.RandomArray(10000)
		InsertionSort(arr)
		if !isIntArraySorted(arr) {
			t.Fatal("Failed")
		}
	}
}
