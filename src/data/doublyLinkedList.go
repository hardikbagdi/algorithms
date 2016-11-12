package data

import "errors"

// DoublyList represents a doubly linked list
type DoublyList struct {
	first, last *Node
	count       int
}

// Front returns the head of the linked list
func (L *DoublyList) Front() *Node {
	return L.first
}

// Tail returns the head of the linked list
func (L *DoublyList) Tail() *Node {
	return L.last
}

// NewDoublyList returns a new empty list
func NewDoublyList() *DoublyList {
	return new(DoublyList)
}

// Size returns the number of nodes in the current linked list
func (L *DoublyList) Size() int {
	return L.count
}

// Empty verifies if the list is empty or not
func (L *DoublyList) Empty() bool {
	return L.count == 0
}

//NodeAt returns the nth(index) node
func (L *DoublyList) NodeAt(index int) *Node {
	if L.count < index {
		return nil
	}
	i := 1
	cur := L.first
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// PushFront inserts an element at the head of the linked list
func (L *DoublyList) PushFront(value Interface) *Node {
	//make new node
	newNode := new(Node)
	newNode.Value = value
	//insert it before first
	newNode.next = L.first
	L.first = newNode
	if L.last == nil {
		L.last = L.first
	}
	L.count++
	return newNode
}

// PopFront removes the head and returns it
func (L *DoublyList) PopFront() *Node {
	if L.first == nil {
		return nil
	}
	node := L.first
	if L.first == L.last {
		L.first, L.last = nil, nil
	} else {
		L.first = L.first.next
	}
	L.count--
	return node
}

// Append adds a node with value val at  the end of the list
func (L *DoublyList) Append(val Interface) *Node {
	if L.count == 0 {
		return L.PushFront(val)
	}
	newNode := new(Node)
	newNode.Value = val
	curLastNode := L.last
	curLastNode.next = newNode
	newNode.prev = curLastNode
	L.last = newNode
	L.count++
	return newNode
}

// PushBack adds the node at the end of the list; same as Append()
func (L *DoublyList) PushBack(value Interface) *Node {
	return L.Append(value)
}

// PopBack removes the last element in the list
func (L *DoublyList) PopBack() *Node {
	if L.count <= 1 {
		return L.PopFront()
	}
	nodeToReturn := L.last
	L.last = L.last.prev
	L.count--
	return nodeToReturn
}

// InsertAt inserts a new nth node with value 'val'
func (L *DoublyList) InsertAt(value Interface, index int) (*Node, error) {
	if index <= 0 || index > L.count {
		return nil, errors.New("Specified index out of range")
	}
	if index == 1 {
		return L.PushFront(value), nil
	}
	newNode := new(Node)
	newNode.Value = value
	i := 1
	prev := L.first
	for ; i < index-1; i++ {
		prev = prev.next
	}
	//save current indexth node
	nextNode := prev.next
	//insert the new node
	prev.next = newNode
	newNode.next = nextNode
	nextNode.prev = newNode
	newNode.prev = prev
	L.count++
	return newNode, nil
}

// RemoveAt removes the nth node (1-based indexing)
func (L *DoublyList) RemoveAt(index int) (*Node, error) {
	if index <= 0 || index > L.count {
		return nil, errors.New("Specified index out of range")
	}
	// we can be certain that these two will not return a nil
	if index == 1 {
		return L.PopFront(), nil
	}
	if index == L.count {
		return L.PopBack(), nil
	}
	prev := L.first
	for i := 1; i < index-1; i++ {
		prev = prev.next
	}
	nodeToReturn := prev.next
	prev.next = prev.next.next
	prev.next.prev = prev
	L.count--
	return nodeToReturn, nil
}
