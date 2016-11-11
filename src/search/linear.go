package search

// LinearSearch searches for searchTerm in the 'array' and returns the index if found
// if searchTerm is not found then -1 is returned
func LinearSearch(array []int, searchTerm int) int {
	if array == nil {
		return -1
	}
	for i, val := range array {
		if val == searchTerm {
			return i
		}
	}
	return -1
}
