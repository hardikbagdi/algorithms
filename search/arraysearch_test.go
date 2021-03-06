package search

import (
	"github.com/hardikbagdi/algorithms/random"
	"github.com/hardikbagdi/algorithms/sorting"
	"math/rand"
	"testing"
)

func testArrayPositiveSearch(T *testing.T) {
	random.RandomizeSeed()
	array := random.Array(100)
	sorting.QuickSort(array)
	searchTerm := array[rand.Intn(len(array))]
	if LinearSearch(array, searchTerm) != BinarySearch(array, searchTerm) {
		T.Fatal("failed")
	}
}

func testArrayNegativeSearch(T *testing.T) {
	random.RandomizeSeed()
	array := random.Array(100)
	sorting.QuickSort(array)
	searchTerm := -42
	x := LinearSearch(array, searchTerm)
	y := BinarySearch(array, searchTerm)
	z := BinarySearchRecursive(array, searchTerm)
	if x != y || x != z {
		T.Fatal("failed")
	}
}

func TestArraySearch(T *testing.T) {
	for i := 0; i < 100; i++ {
		go testArrayNegativeSearch(T)
		go testArrayPositiveSearch(T)
	}
}
