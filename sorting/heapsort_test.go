package sorting

import (
	"random"
	"testing"
)

func testHeapSort(t *testing.T, arr []int) {
	HeapSort(arr)
	if !IsIntArraySorted(arr) {
		t.Fatal("Failed")
	}
}
func TestHeapSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 100; i++ {
		arr := random.Array(10)
		testHeapSort(t, arr)
	}
}
