package linear

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap(T *testing.T) {
	fmt.Println("Heap test enter")
	heap := NewHeap()
	assert.Equal(T, 0, heap.Count())

	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(8)
	heap.Insert(2)
	ele := 0
	ele, _ = heap.ElementAt(0)
	assert.Equal(T, 2, ele)
	fmt.Println(heap.Array())

	index, ok := heap.Search(8)
	assert.Equal(T, 2, index)
	assert.Equal(T, true, ok)

	fmt.Println("decrease key called")
	heap.DecreaseKey(8, 1)
	fmt.Println(heap)

	ele, _ = heap.ElementAt(0)
	assert.Equal(T, 1, ele)

	err := heap.DeleteRoot()
	assert.Nil(T, err)

	ele, _ = heap.ElementAt(0)
	assert.Equal(T, 2, ele)

	index, ok = heap.Search(8)
	assert.Equal(T, 0, index)
	assert.Equal(T, false, ok)

	fmt.Println(heap)
	fmt.Println("Heap test finished")
}
