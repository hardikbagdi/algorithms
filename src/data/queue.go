package data

import "errors"

// Queue represent a queue of items
type Queue struct {
	list *DoublyList
}

// NewQueue returns a new empty queue
func NewQueue() *Queue {
	q := new(Queue)
	q.list = NewDoublyList()
	return q
}

//Enqueue adds the an element to the end of the list
func (Q *Queue) Enqueue(value Interface) {
	Q.list.PushBack(value)
	return
}

// Dequeue removes and returns the first element in the queue
func (Q *Queue) Dequeue() (Interface, error) {
	node := Q.list.PopFront()
	if node == nil {
		return nil, errors.New("Queue underflow")
	}
	return node.Value, nil
}

// Size return the number of elements present in the Queue
func (Q *Queue) Size() int {
	return Q.list.Size()
}
