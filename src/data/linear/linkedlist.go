// Package data contains various data structures
package linear

import "errors"

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
func (L *List) InsertAt(val Interface, index int) (*Node, error) {
	if index <= 0 || index > L.count {
		return nil, errors.New("Specified index out of range")
	}
	if index == 1 {
		return L.PushFront(val), nil
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
	return newNode, nil
}

// RemoveAt removes the nth node (1-based indexing)
func (L *List) RemoveAt(index int) (*Node, error) {
	if index <= 0 || index > L.count {
		return nil, errors.New("Specified index out of range")
	}
	if index == 1 {
		//can be certain that PopFront won't give a nil
		return L.PopFront(), nil
	}
	cur := L.root
	for i := 1; i < index-1; i++ {
		cur = cur.next
	}
	nodeRemoved := cur.next
	cur.next = cur.next.next
	L.count--
	return nodeRemoved, nil
}

// Front returns the head of the linked list
func (L *List) Front() *Node {
	return L.NodeAt(1)
}
