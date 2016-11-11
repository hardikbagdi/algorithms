package sorting

import (
	"random"
	"testing"
)

func testMergeSort(t *testing.T, arr []int) {
	MergeSort(arr)
	if !IsIntArraySorted(arr) {
		t.Fatal("Failed")
	}
}
func TestMergeSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 100; i++ {
		arr := random.RandomArray(100)
		testMergeSort(t, arr)
	}
}
