package sorting

import (
	"github.com/hardikbagdi/algorithms/random"
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
		arr := random.Array(100)
		testMergeSort(t, arr)
	}
}
