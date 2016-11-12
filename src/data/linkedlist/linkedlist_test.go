package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(T *testing.T) {
	list := New()
	assert.Equal(T, 0, list.Size())
	assert.Equal(T, true, list.Empty())
	root := list.Front()
	assert.Nil(T, root)
	assert.Nil(T, list.PopFront())
	assert.Nil(T, list.NodeAt(1))
	assert.Nil(T, list.RemoveAt(1))
	node := list.PushFront(1)
	node1 := new(Node)
	node1.Value = 1
	assert.Equal(T, node, node1)
	assert.Equal(T, 1, list.Size())
	assert.Equal(T, node1, list.NodeAt(1))
	node = list.Append(2)
	node2 := new(Node)
	node2.Value = 2
	assert.Equal(T, node2, list.NodeAt(2))
	node = list.RemoveAt(2)
	assert.Equal(T, node2, node)
	list.Append(3)
	list.Append(4)
	assert.Equal(T, 4, list.NodeAt(3).Value)
	assert.Equal(T, 3, list.Size())
	list.PushFront(42)
	node = list.InsertAt(43, 2)
	assert.Equal(T, node, list.NodeAt(2))
	assert.Equal(T, 4, list.RemoveAt(5).Value)
	node = list.RemoveAt(4)
	node = list.PopFront()
	assert.Equal(T, 42, node.Value)
	list.InsertAt(1, 1)

	list.InsertAt(142, 3)
	assert.Equal(T, 4, list.Size())
	assert.Nil(T, list.InsertAt(442, 442))
	node = list.RemoveAt(1)
	assert.Equal(T, 1, node.Value)
}
