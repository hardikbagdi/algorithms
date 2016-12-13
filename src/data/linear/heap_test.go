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
	heap.Insert(1)

	fmt.Println(heap.Array())

	x := heap.Array()
	fmt.Println(x)
	fmt.Println("Heap test finished")
}
