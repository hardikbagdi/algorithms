package tree

type Comparable interface {
	Compare(Comparable) int
}

type Node struct {
	Value       Comparable
	left, right *Node
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
