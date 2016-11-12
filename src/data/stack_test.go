package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(T *testing.T) {
	stack := NewStack()
	assert.Equal(T, 0, stack.Size())
	stack.Push(1)
	stack.Push(2)
	assert.Equal(T, 2, stack.Size())
	v, _ := stack.Pop()
	assert.Equal(T, 2, v)
	stack.Pop()
	v, err := stack.Pop()
	assert.Nil(T, v)
	assert.NotNil(T, err)
}
