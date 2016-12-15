package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree(T *testing.T) {
	fmt.Println("Tree testing start")
	t := new(Tree)
	assert.NotNil(T, t)
	assert.Nil(T, t.Root())
	root := new(Node)
	root.SetValue(1)
	t.SetRoot(root)
	assert.NotNil(T, t.Root())
	var n2, n3, n4, n5, n6, n7, n8, n9 Node
	n2.SetValue(2)
	n3.SetValue(3)
	n4.SetValue(4)
	n5.SetValue(5)
	n6.SetValue(6)
	n7.SetValue(7)
	n8.SetValue(8)
	n9.SetValue(9)
	root.SetLeft(&n2)
	root.SetRight(&n3)
	n2.SetLeft(&n4)
	n2.SetRight(&n5)
	n5.SetLeft(&n6)
	n5.SetRight(&n7)
	n7.SetLeft(&n8)
	n7.SetRight(&n9)

	preOrder := []int{1, 2, 4, 5, 6, 7, 8, 9, 3}
	inOrder := []int{4, 2, 6, 5, 8, 7, 9, 1, 3}
	postOrder := []int{4, 6, 8, 9, 7, 5, 2, 3, 1}
	i := 0
	for value := range t.PreOrderIterator() {
		//fmt.Println(value.Value())
		assert.Equal(T, preOrder[i], value.Value())
		i++
	}
	i = 0
	for value := range t.InOrderIterator() {
		//fmt.Println(value.Value())
		assert.Equal(T, inOrder[i], value.Value())
		i++
	}
	i = 0
	for value := range t.PostOrderIterator() {
		//fmt.Println(value.Value())
		assert.Equal(T, postOrder[i], value.Value())
		i++
	}
	fmt.Println("BST testing end")
}
