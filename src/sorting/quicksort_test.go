package sorting

import (
	"random"
	"testing"
)

func testQuickSort(t *testing.T, arr []int) {
	QuickSort(arr)
	if !IsIntArraySorted(arr) {
		t.Fatal("Failed")
	}
}
func TestQuickSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 100; i++ {
		arr := random.RandomArray(100)
		testQuickSort(t, arr)
	}
}
