package sorting

import (
	"random"
	"testing"
)

func testQuickSort3way(t *testing.T, arr []int) {
	QuickSort3way(arr)
	if !isIntArraySorted(arr) {
		t.Fatal("Failed")
	}
}
func TestQuickSort3way(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 100; i++ {
		arr := random.RandomArray(100)
		testQuickSort3way(t, arr)
	}
}
