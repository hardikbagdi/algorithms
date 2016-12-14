package linear

import "errors"

// Dequeue represent a deck of items
type Dequeue struct {
	list *DoublyList
}

// NewDequeue returns a new empty dequeue
func NewDequeue() *Dequeue {
	q := new(Dequeue)
	q.list = NewDoublyList()
	return q
}

// AddLast adds the an element to the end of the list
func (Q *Dequeue) AddLast(value Interface) {
	Q.list.PushBack(value)
	return
}

// AddFirst adds the an element to the front of the list
func (Q *Dequeue) AddFirst(value Interface) {
	Q.list.PushFront(value)
	return
}

// PollFirst removes and returns an element from the front of dequeue
func (Q *Dequeue) PollFirst() (Interface, error) {
	node := Q.list.PopFront()
	if node == nil {
		return nil, errors.New("Queue underflow")
	}
	return node.Value, nil
}

// PollLast removes and returns an element from the back of the dequeue
func (Q *Dequeue) PollLast() (Interface, error) {
	node := Q.list.PopBack()
	if node == nil {
		return nil, errors.New("Queue underflow")
	}
	return node.Value, nil
}

// Size return the number of elements present in the Queue
func (Q *Dequeue) Size() int {
	return Q.list.Size()
}
