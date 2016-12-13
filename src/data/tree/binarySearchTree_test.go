package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST(T *testing.T) {
	fmt.Println("BST testing start")
	bst := NewBST()
	assert.Equal(T, 0, bst.Count())

	fmt.Println(bst)
	bst.Insert(3)
	bst.Insert(-33)
	assert.Equal(T, 2, bst.Count())
	bst.Insert(1)
	bst.Insert(10)
	bst.Insert(5)
	fmt.Println(bst)
	assert.Equal(T, true, bst.Search(1))
	assert.Equal(T, true, bst.Search(3))
	assert.Equal(T, false, bst.Search(0))
	err := bst.Remove(3)
	assert.Nil(T, err)
	fmt.Println(bst)
	fmt.Println("BST testing end")
}
