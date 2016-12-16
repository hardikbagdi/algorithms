package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func printHeap(h *Heap) {
	fmt.Println()
	for val := range h.PreOrderIterator() {
		fmt.Printf("%d ", val.Value())
	}
	fmt.Println()
}
func TestHeap(T *testing.T) {
	fmt.Println("Heap testing start")
	h := NewHeap()
	assert.NotNil(T, h)
	assert.Nil(T, h.Root())
	h.Insert(42)
	assert.Equal(T, 1, h.Count())
	assert.NotNil(T, h.Root())
	h.Insert(100)
	n := h.Root()
	assert.NotNil(T, n.Left())
	assert.Nil(T, n.Right())
	err := h.Insert(100)
	assert.NotNil(T, err)
	err = h.Insert(101)
	assert.Nil(T, err)
	assert.NotNil(T, n.Right())
	assert.Equal(T, h.Root(), h.Root().Left().Parent())
	assert.Equal(T, h.Root(), h.Root().Right().Parent())
	printHeap(h)

	h.Insert(20)
	assert.Equal(T, 20, h.Root().Value())
	h.Insert(10)
	assert.Equal(T, 10, h.Root().Value())
	printHeap(h)
	err = h.DecreaseKey(10, 1000)
	assert.NotNil(T, err)
	err = h.IncreaseKey(10, 1000)
	assert.Nil(T, err)
	printHeap(h)

	fmt.Println("Heap testing end")
}
func TestIntToBinaryString(T *testing.T) {
	assert.Equal(T, "11", intToBinaryString(3))
	assert.Equal(T, "101", intToBinaryString(5))
	assert.Equal(T, "100", intToBinaryString(4))
	assert.Equal(T, "11001", intToBinaryString(25))
}
