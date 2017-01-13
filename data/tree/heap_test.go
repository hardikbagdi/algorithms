package tree

import (
	"log"
	"github.com/stretchr/testify/assert"
	"testing"
	"io/ioutil"
)

func verifyHeap(node *Node) bool {
	if node != nil {
		if node.Left() != nil && node.Left().Value() < node.Value() {
			return false
		}
		if node.Right() != nil && node.Right().Value() < node.Value() {
			return false
		}
		return verifyHeap(node.Left()) && verifyHeap(node.Right())
	}
	return true
}

func printHeap(h *Heap) {
	log.Println()
	for val := range h.PreOrderIterator() {
		log.Printf("%d ", val.Value())
	}
	log.Println()
}
func TestHeap(T *testing.T) {
	log.SetOutput(ioutil.Discard)
	log.Println("Heap testing start")
	h := NewHeap()
	assert.NotNil(T, h)
	assert.Nil(T, h.Root())
	assert.NotNil(T, h.Tree())
	err := h.DeleteRoot()
	assert.NotNil(T, err)
	h.Insert(42)
	assert.Equal(T, 1, h.Count())
	assert.NotNil(T, h.Root())
	h.Insert(100)
	n := h.Root()
	assert.NotNil(T, n.Left())
	assert.Nil(T, n.Right())
	err = h.Insert(100)
	assert.NotNil(T, err)
	err = h.Insert(101)
	assert.Nil(T, err)
	assert.NotNil(T, n.Right())
	assert.Equal(T, h.Root(), h.Root().Left().Parent())
	assert.Equal(T, h.Root(), h.Root().Right().Parent())
	printHeap(h)
	assert.Equal(T, true, verifyHeap(h.Root()))

	h.Insert(20)
	assert.Equal(T, 20, h.Root().Value())
	h.Insert(10)
	assert.Equal(T, 10, h.Root().Value())
	printHeap(h)
	err = h.DecreaseKey(10, 1000)
	assert.NotNil(T, err)
	err = h.IncreaseKey(10, 1000)
	assert.Nil(T, err)
	log.Println("pre order")
	printHeap(h)
	assert.Equal(T, true, verifyHeap(h.Root()))
	assert.Equal(T, 5, h.Count())

	inOrder := []int{100, 42, 1000, 20, 101}
	postOrder := []int{100, 1000, 42, 101, 20}
	i := 0
	log.Println("iterators")
	log.Println()
	log.Println("pre order")
	printHeap(h)
	log.Println()
	log.Println("inorder")
	for val := range h.InOrderIterator() {
		log.Printf("%d ", val.Value())
		assert.Equal(T, inOrder[i], val.Value())
		i++
	}
	i = 0
	log.Println()
	log.Println("post-order")
	for val := range h.PostOrderIterator() {
		log.Printf("%d ", val.Value())
		assert.Equal(T, postOrder[i], val.Value())
		i++
	}
	log.Println()
	//postOrder := []int{}

	err = h.IncreaseKey(10000, 12)
	assert.NotNil(T, err)
	err = h.IncreaseKey(-20, 20)
	assert.NotNil(T, err)
	err = h.DecreaseKey(100, 100)
	assert.NotNil(T, err)

	log.Println("deleting root")
	err = h.DeleteRoot()
	assert.Nil(T, err)
	assert.Equal(T, true, verifyHeap(h.Root()))
	assert.Equal(T, 4, h.Count())
	printHeap(h)
	log.Println("deleting 1000")
	err = h.Remove(1000)
	assert.Nil(T, err)
	assert.Equal(T, true, verifyHeap(h.Root()))
	assert.Equal(T, 3, h.Count())
	printHeap(h)
	log.Println("deleting 101")
	err = h.Remove(101)
	assert.Nil(T, err)
	assert.Equal(T, true, verifyHeap(h.Root()))
	assert.Equal(T, 2, h.Count())
	printHeap(h)
	log.Println("deleting 100")
	err = h.Remove(100)
	assert.Nil(T, err)
	assert.Equal(T, true, verifyHeap(h.Root()))
	assert.Equal(T, 1, h.Count())
	printHeap(h)
	err = h.Remove(4242)
	assert.NotNil(T, err)
	log.Println("deleting root")
	err = h.DeleteRoot()
	assert.Nil(T, err)
	assert.Equal(T, true, verifyHeap(h.Root()))
	assert.Equal(T, 0, h.Count())
	assert.Nil(T, h.Root())
	printHeap(h)

	h.Insert(4242)
	printHeap(h)
	assert.Equal(T, 1, h.Count())
	assert.NotNil(T, h.Root())
	assert.Nil(T, h.Root().Left())
	assert.Nil(T, h.Root().Right())
	h.Insert(100)
	assert.Nil(T, h.Root().Right())
	assert.Equal(T, 100, h.Root().Value())
	err = h.DecreaseKey(101, 40)
	assert.NotNil(T, err)
	err = h.DecreaseKey(4242, 40)
	assert.Nil(T, err)
	printHeap(h)
	log.Println("Heap testing end")
}
func TestIntToBinaryString(T *testing.T) {
	assert.Equal(T, "11", intToBinaryString(3))
	assert.Equal(T, "101", intToBinaryString(5))
	assert.Equal(T, "100", intToBinaryString(4))
	assert.Equal(T, "11001", intToBinaryString(25))
}
