package sorting

import (
	"random"
	"testing"
)

func TestMergeSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 10; i++ {
		arr := random.RandomArray(10000)
		MergeSort(arr)
		if !isIntArraySorted(arr) {
			t.Fatal("Failed")
		}
	}
}
