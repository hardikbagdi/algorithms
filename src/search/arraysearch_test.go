package search

import (
	"math/rand"
	"random"
	"sorting"
	"testing"
)

func testArrayPositiveSearch(T *testing.T) {
	random.RandomizeSeed()
	array := random.RandomArray(100)
	sorting.QuickSort(array)
	searchTerm := array[rand.Intn(len(array))]
	if LinearSearch(array, searchTerm) != BinarySearch(array, searchTerm) {
		T.Fatal("failed")
	}
}

func testArrayNegativeSearch(T *testing.T) {
	random.RandomizeSeed()
	array := random.RandomArray(100)
	sorting.QuickSort(array)
	searchTerm := -42
	if LinearSearch(array, searchTerm) != BinarySearch(array, searchTerm) {
		T.Fatal("failed")
	}
}

func TestArraySearch(T *testing.T) {
	for i := 0; i < 100; i++ {
		go testArrayNegativeSearch(T)
		go testArrayPositiveSearch(T)
	}
}
