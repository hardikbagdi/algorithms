package sorting

import (
	"random"
	"testing"
)

func testRadixSort(t *testing.T, arr []int) {
	RadixSort(arr)
	if !IsIntArraySorted(arr) {
		t.Fatal("Failed")
	}
}
func TestRadixSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 100; i++ {
		arr := random.Array(1090)
		testRadixSort(t, arr)
	}
}
