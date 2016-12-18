package sorting

import "sort"
import "fmt"
import "math"

//ArraysEqual returns true if the two int arrays are equal; false otherwise
func ArraysEqual(arr1, arr2 []int) bool {

	for i, val := range arr1 {
		if val != arr2[i] {
			return false
		}
	}
	return true
}

//IsIntArraySorted checks if the given array is sorted or not
func IsIntArraySorted(arr []int) bool {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sort.Ints(arrCopy)
	return ArraysEqual(arr, arrCopy)
}

//Xchg swaps two elements give by index a and b in the array
func Xchg(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

//PrintArray prints the array to the STDOUT
func PrintArray(array []int) {
	fmt.Println("Array: ", array, " length: ", len(array))
}

//GetMax returns the largest number of the given array
func GetMax(array []int) (max int) {

	max = math.MinInt64
	for _, val := range array {
		if max < val {
			max = val
		}
	}
	return
}
