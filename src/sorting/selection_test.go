package sorting

import (
	"random"
	"testing"
)

func testSelectionSort(t *testing.T, arr []int) {
	SelectionSort(arr)
	if !isIntArraySorted(arr) {
		t.Fatal("Failed")
	}
}
func TestSelectionSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 100; i++ {
		arr := random.RandomArray(100)
		testSelectionSort(t, arr)
	}
}
