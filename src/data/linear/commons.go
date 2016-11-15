// Package data contains various data structures
package linear

// Interface to hold data in the nodes
type Interface interface{}

// Node represents one node of a linked list
type Node struct {
	Value      Interface
	next, prev *Node
}

// Next returns the next node
func (N *Node) Next() *Node {
	return N.next
}

// Prev returns the prev pointer for a doubly list
func (N *Node) Prev() *Node {
	return N.prev
}

// SetNext sets the next pointer in the node
func (N *Node) SetNext(nextNode *Node) {
	N.next = nextNode
}

// SetPrev sets the previous pointer in the node
func (N *Node) SetPrev(prevNode *Node) {
	N.prev = prevNode
}
