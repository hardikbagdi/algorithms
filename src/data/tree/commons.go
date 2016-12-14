// Package tree contains tree data structures and operations
package tree

// Node represnts a node of a tree
type Node struct {
	value       int
	left, right *Node
}

// Value returns the integer stored in the Node
func (N *Node) Value() int {
	return N.value
}

// SetValue sets the Value stored in the Node
func (N *Node) SetValue(val int) {
	N.value = val
}

// Left returns the pointer to the left child
func (N *Node) Left() *Node {
	return N.left
}

// Right returns the pointer to the right child
func (N *Node) Right() *Node {
	return N.right
}

// SetLeft sets the left child of the Node
func (N *Node) SetLeft(leftNode *Node) {
	N.left = leftNode
}

// SetRight sets the right child of the Node
func (N *Node) SetRight(rightNode *Node) {
	N.right = rightNode
}

// Tree represents a binary tree data type
type Tree struct {
	root *Node
}

// Root returns the root node of the tree
func (T *Tree) Root() *Node {
	return T.root
}
