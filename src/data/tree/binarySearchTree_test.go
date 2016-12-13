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
	fmt.Println(bst)
	fmt.Println("BST testing end")
}
