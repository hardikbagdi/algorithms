package sorting

import (
	"github.com/hardikbagdi/algorithms/random"
	"testing"
)

func testInsertionSort(t *testing.T, arr []int) {
	InsertionSort(arr)
	if !IsIntArraySorted(arr) {
		t.Fatal("Failed")
	}
}
func TestInsertionSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 100; i++ {
		arr := random.Array(100)
		testInsertionSort(t, arr)
	}
}
