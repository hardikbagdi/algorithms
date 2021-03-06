package main

import (
	"fmt"
	"github.com/hardikbagdi/algorithms/data/linear"
	"github.com/hardikbagdi/algorithms/random"
	"github.com/hardikbagdi/algorithms/search"
	"github.com/hardikbagdi/algorithms/sorting"
	"math/rand"
)

func main() {
	const (
		arraySize  = 20
		iterations = 10
	)
	random.RandomizeSeed()
	for i := 0; i < iterations; i++ {
		x := random.Array(arraySize)
		fmt.Println("Before sorting")
		fmt.Println(x)
		fmt.Println("After sorting")
		//sorting.BubbleSort(x)
		//sorting.InsertionSort(x)
		//sorting.SelectionSort(x)
		//sorting.MergeSort(x)
		//sorting.QuickSort(x)
		//sorting.QuickSort3way(x)
		sorting.RadixSort(x)
		fmt.Println(x)
		searchTerm := x[rand.Intn(arraySize)]
		fmt.Println(search.LinearSearch(x, searchTerm))
		fmt.Println(search.BinarySearch(x, searchTerm))
	}
	stack := linear.NewStack()
	stack.Push(1)
}
