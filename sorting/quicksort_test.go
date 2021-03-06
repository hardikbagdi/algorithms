package sorting

import (
	"github.com/hardikbagdi/algorithms/random"
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
		arr := random.Array(100)
		testQuickSort(t, arr)
	}
}
