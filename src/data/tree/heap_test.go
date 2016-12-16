package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap(T *testing.T) {
	fmt.Println("Heap testing start")
	h := NewHeap()
	assert.NotNil(T, h)
	assert.Nil(T, h.Root())
	fmt.Println("Heap testing end")
}
