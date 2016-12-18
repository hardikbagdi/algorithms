package tree

import (
	"bytes"
	"errors"
	"github.com/hardikbagdi/algorithms/data/linear"
	"strconv"
)

// BST represents a binary search tree
type BST struct {
	tree  *Tree
	count int
}

// NewBST returns a new BST
func NewBST() *BST {
	bst := new(BST)
	bst.tree = new(Tree)
	bst.count = 0
	return bst
}

// Tree returns the backing binary tree
func (bst *BST) Tree() *Tree {
	return bst.tree
}

// String returns string representation of the BST
func (bst *BST) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Count: ")
	buffer.WriteString(strconv.Itoa(bst.count) + "\n")
	stack := linear.NewStack()

	current := bst.Tree().Root()
	done := false
	for !done {
		if current != nil {
			stack.Push(current)
			current = current.Left()
		} else {
			if stack.Size() != 0 {
				temp, _ := stack.Pop()
				current = temp.(*Node)
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
	newRoot := bst.insertHelper(bst.Tree().Root(), value)
	bst.Tree().SetRoot(newRoot)
	bst.count++
	return nil
}

func (bst *BST) getNextPreOrder(node *Node) int {
	key := node.Value()
	for node.Left() != nil {
		node = node.Left()
		key = node.Value()
	}
	return key
}

func (bst *BST) removeHelper(node *Node, value int) *Node {
	if node == nil {
		return node
	}
	key := node.Value()
	if value < key {
		node.SetLeft(bst.removeHelper(node.Left(), value))
	} else if key < value {
		node.SetRight(bst.removeHelper(node.Right(), value))
	} else {
		if node.Left() == nil {
			return node.Right()
		} else if node.Right() == nil {
			return node.Left()
		} else {
			//next inorder elment
			key = bst.getNextPreOrder(node.Right())
			node.SetValue(key)
			node.SetRight(bst.removeHelper(node.Right(), key))
		}
	}
	return node
}

// Remove a value from the array
func (bst *BST) Remove(value int) error {
	present := bst.Search(value)
	if !present {
		return errors.New("Value not present in BST")
	}
	newRoot := bst.removeHelper(bst.Tree().Root(), value)
	bst.Tree().SetRoot(newRoot)
	bst.count--
	return nil
}

func (bst *BST) searchHelper(node *Node, value int) bool {
	if node == nil {
		return false
	}
	key := node.Value()
	if value == key {
		return true
	} else if value < key {
		return bst.searchHelper(node.Left(), value)
	} else {
		return bst.searchHelper(node.Right(), value)
	}
}

// Search performs search  on the backing array
func (bst *BST) Search(value int) bool {
	return bst.searchHelper(bst.Tree().Root(), value)
}
