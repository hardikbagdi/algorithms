package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(T *testing.T) {
	queue := NewQueue()
	assert.Equal(T, 0, queue.Size())
	queue.Enqueue(1)
	queue.Enqueue(2)
	assert.Equal(T, 2, queue.Size())
	v, err := queue.Dequeue()
	assert.Nil(T, err)
	assert.Equal(T, 1, v)
	queue.Dequeue()
	v, err = queue.Dequeue()
	assert.Nil(T, v)
	assert.NotNil(T, err)
}
