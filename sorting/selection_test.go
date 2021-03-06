package sorting

import (
	"github.com/hardikbagdi/algorithms/random"
	"testing"
)

func testSelectionSort(t *testing.T, arr []int) {
	SelectionSort(arr)
	if !IsIntArraySorted(arr) {
		t.Fatal("Failed")
	}
}
func TestSelectionSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 100; i++ {
		arr := random.Array(100)
		testSelectionSort(t, arr)
	}
}
