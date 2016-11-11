package main

import "fmt"
import "sorting"
import "random"

func main() {
	random.RandomizeSeed()
	for i := 0; i < 100; i++ {
		x := random.RandomArray(20)
		fmt.Println("Before sorting")
		fmt.Println(x)
		fmt.Println("After sorting")
		//sorting.BubbleSort(x)
		//sorting.InsertionSort(x)
		//sorting.SelectionSort(x)
		//sorting.MergeSort(x)
		//sorting.QuickSort(x)
		sorting.QuickSort3way(x)
		fmt.Println(x)
	}
}
