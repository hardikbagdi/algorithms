//Provides basic sorting functions for 'int' datatype.
//Wish there were generics
package sorting

// Bubble sort an array of int
func BubbleSort(array []int) {
	if array == nil {
		return
	}
	length := len(array)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
