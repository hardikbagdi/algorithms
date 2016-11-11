package search

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
