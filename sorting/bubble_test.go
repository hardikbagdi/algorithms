package sorting

import "testing"
import "random"

func testBubbleSort(t *testing.T, arr []int) {
	BubbleSort(arr)
	if !IsIntArraySorted(arr) {
		t.Fatal("Failed")
	}
}
func TestBubbleSort(t *testing.T) {
	random.RandomizeSeed()
	for i := 0; i < 100; i++ {
		arr := random.Array(100)
		testBubbleSort(t, arr)
	}
}
