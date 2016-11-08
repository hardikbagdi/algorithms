package main

import "fmt"
import "sorting"

func main() {
	x := []int{12, 23, 12, 5, 66, 35, 2352, 36, 3}
	sorting.BubbleSort(x)
	fmt.Println(x)
}
