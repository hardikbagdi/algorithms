package sorting

import (
	"random"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 10; i++ {
		arr := random.RandomArray(10000)
		SelectionSort(arr)
		if !isIntArraySorted(arr) {
			t.Fatal("Failed")
		}
	}
}
