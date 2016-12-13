package tree

type Node struct {
	value       int
	left, right *Node
}

func (N *Node) Value() int {
	return N.value
}

func (N *Node) SetValue(val int) {
	N.value = val
}

func (N *Node) Left() *Node {
	return N.left
}

func (N *Node) Right() *Node {
	return N.right
}

func (N *Node) SetLeft(leftNode *Node) {
	N.left = leftNode
}

func (N *Node) SetRight(rightNode *Node) {
	N.right = rightNode
}

type Tree struct {
	root *Node
}

func (T *Tree) Root() *Node {
	return T.root
}
