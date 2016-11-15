package linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(T *testing.T) {
	list := NewLinkedList()
	assert.Equal(T, 0, list.Size())
	assert.Equal(T, true, list.Empty())
	root := list.Front()
	assert.Nil(T, root)
	assert.Nil(T, list.PopFront())
	assert.Nil(T, list.NodeAt(1))
	node, err := list.RemoveAt(1)
	assert.Nil(T, node)
	assert.NotNil(T, err)
	node = list.PushFront(1)
	node1 := new(Node)
	node1.Value = 1
	assert.Equal(T, node, node1)
	assert.Equal(T, 1, list.Size())
	assert.Equal(T, node1, list.NodeAt(1))
	node = list.Append(2)
	node2 := new(Node)
	node2.Value = 2
	assert.Equal(T, node2, list.NodeAt(2))
	node, err = list.RemoveAt(2)
	assert.Equal(T, node2, node)
	list.Append(3)
	list.Append(4)
	assert.Equal(T, 4, list.NodeAt(3).Value)
	assert.Equal(T, 3, list.Size())
	list.PushFront(42)
	node, err = list.InsertAt(43, 2)
	assert.Equal(T, node, list.NodeAt(2))
	node, err = list.RemoveAt(5)
	assert.Equal(T, 4, node.Value)
	node, err = list.RemoveAt(4)
	node = list.PopFront()
	assert.Equal(T, 42, node.Value)
	list.InsertAt(1, 1)

	list.InsertAt(142, 3)
	assert.Equal(T, 4, list.Size())
	node, err = list.InsertAt(442, 442)
	assert.Nil(T, node)
	assert.NotNil(T, err)
	node, err = list.RemoveAt(1)
	assert.Nil(T, err)
	assert.Equal(T, 1, node.Value)
}
