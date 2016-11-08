package sorting

import "fmt"

//func Sum(a,b int) int {
//	return a+b
//}

func BubbleSort(array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func main() {
	arr := []int{12, 4, 325, 2, 66, 7, 7, 3, 35}
	BubbleSort(arr)
	fmt.Println(arr)
}
