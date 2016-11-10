package main

import "fmt"
import "sorting"
import "random"

func main() {
	x := random.RandomArray(10)
	sorting.BubbleSort(x)
	fmt.Println(x)
}
