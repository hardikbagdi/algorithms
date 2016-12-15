package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func verifyBST(node *Node) bool {
	if node == nil {
		return true
	}
	if node.Left() != nil {
		if node.Value() < node.Left().Value() {
			return false
		}
	}
	if node.Right() != nil {
		if node.Value() > node.Right().Value() {
			return false
		}
	}
	return verifyBST(node.Left()) && verifyBST(node.Right())
}

func TestBST(T *testing.T) {
	fmt.Println("BST testing start")
	t := time.Now()
	rand.Seed(t.Unix())
	bst := NewBST()
	assert.Equal(T, 0, bst.Count())

	fmt.Println(bst)
	bst.Insert(3)
	bst.Insert(-33)
	assert.Equal(T, 2, bst.Count())
	bst.Insert(1)
	bst.Insert(10)
	bst.Insert(5)
	verifyBST(bst.Tree().Root())
	fmt.Println(bst)
	assert.Equal(T, true, bst.Search(1))
	assert.Equal(T, true, bst.Search(3))
	assert.Equal(T, false, bst.Search(0))
	err := bst.Remove(3)
	assert.Nil(T, err)

	for i := -20; i < 20; i++ {
		bst.Insert(i)
	}
	verifyBST(bst.Tree().Root())
	for i := -20; i < 20; i++ {
		bst.Remove(i)
	}
	err = bst.Remove(1)
	assert.NotNil(T, err)
	verifyBST(bst.Tree().Root())
	for i := 1; i < 10000; i++ {
		bst.Insert(rand.Intn(10000))
		verifyBST(bst.Tree().Root())
	}

	for i := 1; i < 10000; i++ {
		bst.Remove(rand.Intn(10000))
		verifyBST(bst.Tree().Root())
	}

	fmt.Println(bst)
	fmt.Println("BST testing end")
}
