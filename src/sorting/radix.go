package sorting

// Sort an array of Int using radix sort
func RadixSort(array []int) {
	max := GetMax(array)

	for i := 1; max/i > 0; i *= 10 {
		radixsort(array, i)
	}
}

func radixsort(array []int, digit int) {

	var count [10]int
	//backup array
	arrCopy := make([]int, len(array))

	for _, val := range array {
		count[(val/digit)%10]++
	}

	//add the counts for indexing
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	copy(arrCopy, array)

	for i := len(array) - 1; i >= 0; i-- {
		array[count[(arrCopy[i]/digit)%10]-1] = arrCopy[i]
		count[(arrCopy[i]/digit)%10]--
	}
}
