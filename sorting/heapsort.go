package sorting

// HeapSort sorts an array of Int using heapsort
func HeapSort(array []int) {
	length := len(array)
	for i := (length/2 - 1); i >= 0; i-- {
		heapify(array, i)
	}

	for len := length - 1; len >= 1; len-- {
		array[len], array[0] = array[0], array[len]
		heapify(array[:len], 0)
	}
}

func heapify(array []int, start int) {
	largest := start
	left := start*2 + 1
	right := left + 1

	if left < len(array) && array[left] > array[largest] {
		largest = left
	}
	if right < len(array) && array[right] > array[largest] {
		largest = right
	}
	if largest != start {
		array[largest], array[start] = array[start], array[largest]
		heapify(array, largest)
	}
}
