package graph

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBFSGraph(T *testing.T) {
	fmt.Println("testing BFS start")
	g := New(4)
	var edge Edge
	edge.SetSrc(0)
	edge.SetDest(1)
	g.AddEdge(edge)
	edge.SetDest(2)
	g.AddEdge(edge)
	edge.SetSrc(3)
	g.AddEdge(edge)
	edge.SetDest(3)
	g.AddEdge(edge)
	g.print()
	assert.Equal(T, 4, g.Edges())
	bfsiter := []int{0, 2, 1, 3}
	ch, err := g.BFS(0)
	assert.Nil(T, err)
	i := 0
	for node := range ch {
		assert.Equal(T, bfsiter[i], node)
		i++
	}
	fmt.Println("testing BFS end")
}
