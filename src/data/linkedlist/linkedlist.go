// Package linkedlist contains various data structures
package linkedlist

// Interface to hold data in the nodes
type Interface interface{}

// Node represents one node of a linked list
type Node struct {
	Value Interface
	Next  *Node
}

// List represents a singly linked list
type List struct {
	Root  *Node // Why should root be accessible?
	count int
}

var nullList *List
var nulValue Interface

// New returns a new empty list
func New() *List {
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
	cur := L.Root
	for ; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

// PushFront inserts an element at the head of the linked list
func (L *List) PushFront(value Interface) *Node {
	newNode := new(Node)
	newNode.Value = value
	newNode.Next = L.Root
	L.Root = newNode
	L.count++
	return newNode
}

// PopFront removes the head and returns it
func (L *List) PopFront() *Node {
	if L.Root == nil {
		return nil
	}
	node := L.Root
	L.Root = L.Root.Next
	L.count--
	return node
}

// Append adds a node with value val at  the end of the list
func (L *List) Append(val Interface) *Node {
	newNode := new(Node)
	newNode.Value = val
	cur := L.Root
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
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
	cur := L.Root
	for ; i < index-1; i++ {
		cur = cur.Next
	}
	//save current indexth node
	next := cur.Next
	//insert the new node
	cur.Next = newNode
	newNode.Next = next
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
	cur := L.Root
	i := 1
	for ; i < index-1; i++ {
		cur = cur.Next
	}
	nodeRemoved := cur.Next
	cur.Next = cur.Next.Next
	return nodeRemoved
}
