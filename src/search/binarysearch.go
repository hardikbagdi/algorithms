package search

// BinarySearch looks for searchTerm in the sorted array 'array' and return the index
// array has to be sorted
// if searchTerm is not found then -1 is returned
func BinarySearch(array []int, searchTerm int) int {
	if array == nil {
		return -1
	}

	lo := 0
	high := len(array) - 1
	mid := -1
	for lo <= high {
		mid = lo + (high-lo)/2
		comp := searchTerm - array[mid]
		if comp == 0 { //match
			return mid
		} else if comp > 0 {
			lo = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// BinarySearchRecursive performs search for a term on a sorted array
func BinarySearchRecursive(array []int, searchTerm int) int {
	return binarySearchRecHelper(array, searchTerm, 0, len(array)-1)
}

func binarySearchRecHelper(array []int, searchTerm, lo, high int) int {
	if array == nil || high < lo {
		return -1
	}
	mid := lo + (high-lo)/2
	comp := searchTerm - array[mid]
	if comp == 0 { //match
		return mid
	} else if comp > 0 {
		lo = mid + 1
	} else {
		high = mid - 1
	}
	return binarySearchRecHelper(array, searchTerm, lo, high)
}
