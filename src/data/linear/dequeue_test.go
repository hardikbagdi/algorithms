package linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDequeue(T *testing.T) {
	dequeue := NewDequeue()
	v, err := dequeue.PollFirst()
	assert.NotNil(T, err)
	assert.Equal(T, 0, dequeue.Size())
	dequeue.AddFirst(1)
	dequeue.AddLast(2)
	assert.Equal(T, 2, dequeue.Size())
	v, err = dequeue.PollFirst()
	assert.Nil(T, err)
	assert.Equal(T, 1, v)
	v, err = dequeue.PollLast()
	assert.Nil(T, err)
	assert.Equal(T, 2, v)
	v, err = dequeue.PollLast()
	assert.NotNil(T, err)
	v, err = dequeue.PollFirst()
	assert.NotNil(T, err)
}
