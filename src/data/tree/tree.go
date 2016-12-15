package tree

import (
	"bytes"
	"strconv"
)

// Tree represents a binary tree data type
type Tree struct {
	root *Node
}

// Root returns the root node of the tree
func (T *Tree) Root() *Node {
	return T.root
}

// SetRoot sets the root of the tree
func (T *Tree) SetRoot(node *Node) {
	T.root = node
}

// PreOrder returns pre-order traversal
func (T *Tree) PreOrder() string {
	var buffer bytes.Buffer
	helperPreOrder(T.Root(), buffer)
	return buffer.String()
}

func helperPreOrder(node *Node, buffer bytes.Buffer) {
	buffer.WriteString(strconv.Itoa(node.Value()))
	helperPreOrder(node.Left(), buffer)
	helperPreOrder(node.Right(), buffer)
}

// InOrder returns in-order traversal
func (T *Tree) InOrder() string {
	var buffer bytes.Buffer
	helperInOrder(T.Root(), buffer)
	return buffer.String()
}

func helperInOrder(node *Node, buffer bytes.Buffer) {
	helperInOrder(node.Left(), buffer)
	buffer.WriteString(strconv.Itoa(node.Value()))
	helperInOrder(node.Right(), buffer)
}

// PostOrder returns post-order traversal
func (T *Tree) PostOrder() string {
	var buffer bytes.Buffer
	helperPostOrder(T.Root(), buffer)
	return buffer.String()
}

func helperPostOrder(node *Node, buffer bytes.Buffer) {
	helperPostOrder(node.Left(), buffer)
	helperPostOrder(node.Right(), buffer)
	buffer.WriteString(strconv.Itoa(node.Value()))
}
