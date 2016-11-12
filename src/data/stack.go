package data

import "errors"

// Stack represents a LIFO stack
type Stack struct {
	list *List
}

// NewStack returns a new empty stack
func NewStack() *Stack {
	s := new(Stack)
	s.list = NewLinkedList()
	return s
}

// Pop removes and returns the top element on the stack
func (S *Stack) Pop() (Interface, error) {
	newNode := S.list.PopFront()
	if newNode == nil {
		return nil, errors.New("Stack Underflow")
	}
	return newNode.Value, nil
}

// Push puts 'value' on top of the stack
func (S *Stack) Push(value Interface) {
	S.list.PushFront(value)
	return
}

// Size return the number of elements present in the Stack
func (S *Stack) Size() int {
	return S.list.Size()
}
