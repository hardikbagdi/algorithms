package graph

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testdfs1(assert *assert.Assertions) {
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
	assert.Equal(4, g.Edges())
	bfsiter := []int{0, 1, 2, 3}
	ch, err := g.DFS(0)
	assert.Nil(err)
	i := 0
	for node := range ch {
		assert.Equal(bfsiter[i], node)
		i++
	}
}

func testdfs2(assert *assert.Assertions) {
	g := New(11)
	var edge Edge
	edges := [][]int{{0, 2}, {0, 1}, {0, 3}, {1, 3}, {1, 4}, {3, 4}, {4, 5},
		{6, 7}, {6, 9}, {6, 10}, {7, 10}, {8, 9}}
	log.Println("Number of Edges : ", len(edges))
	for i := 0; i < len(edges); i++ {
		edge.SetSrc(edges[i][0])
		edge.SetDest(edges[i][1])
		err := g.AddEdge(edge)
		assert.Nil(err)
	}
	assert.Equal(len(edges), g.Edges())
	g.print()
	bfsiter := []int{0, 2, 1, 3, 4, 5, 6, 7, 10, 9, 8}
	ch, err := g.DFS(0)
	assert.Nil(err)
	i := 0
	for node := range ch {
		assert.Equal(bfsiter[i], node)
		i++
	}
	bfsiter = []int{10, 6, 7, 9, 8, 0, 2, 1, 3, 4, 5}
	ch, err = g.DFS(10)
	assert.Nil(err)
	i = 0
	for node := range ch {
		assert.Equal(bfsiter[i], node)
		i++
	}
}

func TestDFSGraph(T *testing.T) {
	//log.SetOutput(ioutil.Discard)
	assert := assert.New(T)
	log.Println("testing BFS start")
	testdfs1(assert)
	testdfs2(assert)
	log.Println("testing BFS end")
}
