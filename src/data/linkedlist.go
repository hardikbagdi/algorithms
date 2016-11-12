// Package data contains various data structures
package data

// Node represents one node of a linked list
type Node struct {
	Value Interface
	next  *Node
}

// Next returns the next node
func (N *Node) Next() *Node {
	return N.next
}

// SetNext sets the next pointer in the node
func (N *Node) SetNext(nextNode *Node) {
	N.next = nextNode
}

// List represents a singly linked list
type List struct {
	root  *Node // Why should root be accessible?
	count int
}

var nullList *List
var nulValue Interface

// NewLinkedList returns a new empty list
func NewLinkedList() *List {
	return new(List)
}

// Size returns the number of nodes in the current linked list
func (L *List) Size() int {
	return L.count
}

// Empty verifies if the list is empty or not
func (L *List) Empty() bool {
	return L.count == 0
}

//NodeAt returns the nth(index) node
func (L *List) NodeAt(index int) *Node {
	if L.count < index {
		return nil
	}
	i := 1
	cur := L.root
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// PushFront inserts an element at the head of the linked list
func (L *List) PushFront(value Interface) *Node {
	newNode := new(Node)
	newNode.Value = value
	newNode.next = L.root
	L.root = newNode
	L.count++
	return newNode
}

// PopFront removes the head and returns it
func (L *List) PopFront() *Node {
	if L.root == nil {
		return nil
	}
	node := L.root
	L.root = L.root.next
	L.count--
	return node
}

// Append adds a node with value val at  the end of the list
func (L *List) Append(val Interface) *Node {
	newNode := new(Node)
	newNode.Value = val
	cur := L.root
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
	L.count++
	return newNode
}

// InsertAt inserts a new nth node with value 'val'
func (L *List) InsertAt(val Interface, index int) *Node {
	if index <= 0 || index > L.count {
		return nil
	}
	if index == 1 {
		return L.PushFront(val)
	}
	newNode := new(Node)
	newNode.Value = val
	i := 1
	cur := L.root
	for ; i < index-1; i++ {
		cur = cur.next
	}
	//save current indexth node
	nextNode := cur.next
	//insert the new node
	cur.next = newNode
	newNode.next = nextNode
	L.count++
	return newNode
}

// RemoveAt removes the nth node (1-based indexing)
func (L *List) RemoveAt(index int) *Node {
	if index <= 0 || index > L.count {
		return nil
	}
	if index == 1 {
		return L.PopFront()
	}
	cur := L.root
	i := 1
	for ; i < index-1; i++ {
		cur = cur.next
	}
	nodeRemoved := cur.next
	cur.next = cur.next.next
	L.count--
	return nodeRemoved
}

// Front returns the head of the linked list
func (L *List) Front() *Node {
	return L.NodeAt(1)
}
