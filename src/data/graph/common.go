// Package graph has all graph data structures
package graph

import (
	"data/linear"

	"errors"
	"fmt"
)

// Edge represents and edge in a Graph
type Edge struct {
	src, dst, weight int
}

// Src return the source of an edge if the edge is directed.
// For an undirected edge, it's one of the endpoints.
func (E *Edge) Src() int {
	return E.src
}

// Dest return the destination of an edge if the edge is directed.
// For an undirected edge, it's one of the endpoints.
func (E *Edge) Dest() int {
	return E.dst
}

// Weight returns the weight of an edge.
// For an unweighted graph, this should be ignored.
func (E *Edge) Weight() int {
	return E.weight
}

// SetSrc sets the source node of an edge.
func (E *Edge) SetSrc(newSrc int) {
	E.src = newSrc
}

// SetDest sets the destination node of an edge.
// For an undirected graph, this sets another end point
func (E *Edge) SetDest(newDest int) {
	E.dst = newDest
}

// SetWeight sets the weight of an edge.
func (E *Edge) SetWeight(newWeight int) {
	E.weight = newWeight
}

// Graph represents an undirected graph
type Graph struct {
	v, e    int
	adjList []*linear.List
}

// New returns a new undirected Graph
func New(nodes int) *Graph {
	g := new(Graph)
	g.adjList = make([]*linear.List, nodes)
	for i := range g.adjList {
		g.adjList[i] = linear.NewLinkedList()
	}
	g.v = nodes
	return g
}

// Nodes return the number of nodes in the graph.
func (G *Graph) Nodes() int {
	return G.v
}

// Edges returns the number of edges in the graph.
func (G *Graph) Edges() int {
	return G.e
}

// AddEdge adds an undirected edge in the graph.
func (G *Graph) AddEdge(edge Edge) error {
	if !G.validEdge(edge) {
		return errors.New("The given edge is not valid")
	}
	if G.Present(edge) {
		return errors.New("The given edge is already present in the graph")
	}
	edge2 := G.reverseEdge(edge)
	srcList := G.adjList[edge.Src()]
	destList := G.adjList[edge.Dest()]
	srcList.PushFront(edge)
	destList.PushFront(edge2)
	G.e++
	return nil
}

// DeleteEdge deletes an undirected edge in the graph.
func (G *Graph) DeleteEdge(edge Edge) error {
	if !G.validEdge(edge) {
		return errors.New("The given edge is not valid")
	}
	if !G.Present(edge) {
		return errors.New("The edge to delete is not  present in the graph")
	}
	edge2 := G.reverseEdge(edge)
	srcList := G.adjList[edge.Src()]
	destList := G.adjList[edge.Dest()]
	G.deleteHelper(*srcList, edge)
	G.deleteHelper(*destList, edge2)
	G.e--
	return nil
}

func (G *Graph) deleteHelper(list linear.List, edgeToDelete Edge) {
	i := 1
	ch, _ := G.AdjEdges(edgeToDelete.Src())
	for edge := range ch {
		if edgeToDelete.Dest() == edge.Dest() {
			list.RemoveAt(i)
			return
		}
		i++
	}
}

// Present returns true if edgeToFind is present in the graph.
func (G *Graph) Present(edgeToFind Edge) bool {
	if !G.validEdge(edgeToFind) {
		return false
	}
	ch, _ := G.AdjEdges(edgeToFind.Src()) //can ignore error as validEdge above takes care
	for edge := range ch {
		if edgeToFind.Dest() == edge.Dest() {
			return true
		}
	}
	return false
}

// AdjEdges returns a channel of Edge which are edges connected with the source.
// error will be set if the source is not valid
func (G *Graph) AdjEdges(source int) (chan Edge, error) {
	if !G.validNode(source) {
		return nil, errors.New("The given node is not valid")
	}
	ch := make(chan Edge)
	list := G.adjList[source]
	head := list.Front()
	go func(head *linear.Node, ch chan Edge) {
		cur := head
		for cur != nil {
			ch <- cur.Value.(Edge)
			cur = cur.Next()
		}
		close(ch)
	}(head, ch)
	return ch, nil
}
func (G *Graph) printGraph() {
	fmt.Println("Graph:")
	fmt.Println("Nodes: ", G.Nodes(), ", Edges: ", G.Edges())
	for i := range G.adjList {
		fmt.Println("Node: ", i, "\nEdges: ")
		ch, _ := G.AdjEdges(i)
		for edge := range ch {
			fmt.Printf(" (%d,%d)", edge.Src(), edge.Dest())
		}
	}
}
func (G *Graph) validEdge(edge Edge) bool {
	if G.validNode(edge.Src()) && G.validNode(edge.Dest()) {
		return true
	}
	return false
}
func (G *Graph) validNode(node int) bool {
	if node < 0 || node >= G.Nodes() {
		return false
	}
	return true
}

func (G *Graph) reverseEdge(edge Edge) Edge {
	edge2 := edge
	edge2.SetSrc(edge.Dest())
	edge2.SetDest(edge.Src())
	return edge2
}
