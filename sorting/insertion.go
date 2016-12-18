package sorting

// InsertionSort  sorts an array of int
func InsertionSort(array []int) {
	if array == nil {
		return
	}
	length := len(array)
	for i := 1; i < length; i++ {
		key := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > key {
				array[j+1] = array[j]
			} else {
				break
			}
		}
		array[j+1] = key
	}
}
