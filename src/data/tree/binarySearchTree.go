package tree

import (
	"bytes"
	"errors"
	"linear"
	"strconv"
)

// BST represents a binary search tree
type BST struct {
	root  *Node
	count int
}

// NewBST returns a new BST
func NewBST() *BST {
	bst := new(BST)
	bst.count = 0
	return bst
}

// String returns string representation of the BST
func (bst *BST) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Count: ")
	buffer.WriteString(strconv.Itoa(bst.count) + "\n")
	stack := linear.NewStack()

	current := bst.root
	done := false
	for !done {
		if current != nil {
			stack.Push(current)
			current = current.Left()
		} else {
			if stack.Size != 0 {
				current = stack.Pop()
				buffer.WriteString(" " + strconv.Itoa(current.Value()))
				current = current.Right()
			} else {
				done = true
			}
		}

	}
	return buffer.String()
}

// Count returns the number of elements currently in the bst
func (bst *BST) Count() int {
	return bst.count
}

func (bst *BST) insertHelper(node *Node, value int) *Node {
	if node == nil {
		node = new(Node)
		node.SetValue(value)
		return node
	}
	key := node.Value()
	if value < key {
		node.SetLeft(bst.insertHelper(node.Left(), value))
	} else {
		node.SetRight(bst.insertHelper(node.Right(), value))
	}
	return node

}

// Insert inserts a new value into the BST
func (bst *BST) Insert(value int) error {
	present := bst.Search(value)
	if present {
		return errors.New("Value already in BST")
	}
	bst.root = bst.insertHelper(bst.root, value)
	bst.count++
	return nil
}

func (bst *BST) removeHelper(node *Node, value int) *Node {

	return nil
}

// Remove a value from the array
func (bst *BST) Remove(value int) error {
	present := bst.Search(value)
	if !present {
		return errors.New("Value not present in BST")
	}
	bst.root = bst.removeHelper(bst.root, value)
	return nil
}

func (bst *BST) searchHelper(node *Node, value int) bool {
	if node == nil {
		return false
	}
	key := node.Value()
	if value < key {
		return bst.searchHelper(node.Left(), value)
	} else {
		return bst.searchHelper(node.Left(), value)
	}
}

// Search performs search  on the backing array
func (bst *BST) Search(value int) bool {
	return bst.searchHelper(bst.root, value)
}
