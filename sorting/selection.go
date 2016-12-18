package sorting

// SelectionSort sorts an array of int
func SelectionSort(array []int) {
	if array == nil {
		return
	}
	length := len(array)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		array[minIndex], array[i] = array[i], array[minIndex]
	}
}
