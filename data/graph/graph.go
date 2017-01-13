// Package graph has all graph data structures
package graph

import (
	"github.com/hardikbagdi/algorithms/data/linear"

	"errors"
	"log"
)

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
	if edge.Src() != edge.Dest() {
		edge2 := G.reverseEdge(edge)
		destList := G.adjList[edge.Dest()]
		destList.PushFront(edge2)
	}
	srcList := G.adjList[edge.Src()]
	srcList.PushFront(edge)
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
	if edge.Src() != edge.Dest() {
		edge2 := G.reverseEdge(edge)
		destList := G.adjList[edge.Dest()]
		G.deleteHelper(destList, edge2)
	}
	srcList := G.adjList[edge.Src()]
	G.deleteHelper(srcList, edge)
	G.e--
	return nil
}

func (G *Graph) deleteHelper(list *linear.List, edgeToDelete Edge) {
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
	// TODO should check in edge.Dest() list as well?
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
func (G *Graph) print() {
	log.Println("Graph:")
	log.Println("Nodes: ", G.Nodes(), ", Edges: ", G.Edges())
	for i := range G.adjList {
		log.Println("\nNode: ", i, "\nEdges: ")
		ch, _ := G.AdjEdges(i)
		for edge := range ch {
			log.Printf(" (%d,%d)", edge.Src(), edge.Dest())
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

// BFS returns a breadth-first search iteration channel from a node
func (G *Graph) BFS(source int) (chan int, error) {
	if !G.validNode(source) {
		return nil, errors.New("The source is not a valid node in the graph")
	}
	ch := make(chan int, 20)
	go func(source int) {
		visited := make([]bool, G.Nodes())
		G.bfsHelper(source, visited, ch)
		for i, value := range visited {
			if !value {
				G.bfsHelper(i, visited, ch)
			}
		}
		close(ch)
	}(source)
	return ch, nil
}

func (G *Graph) bfsHelper(source int, visited []bool, ch chan int) {
	q := linear.NewQueue()
	q.Enqueue(source)
	visited[source] = true
	for q.Size() != 0 {
		cur, _ := q.Dequeue()
		node, _ := cur.(int)
		ch <- node
		log.Println("BFS EDGE: ", node)
		edges, _ := G.AdjEdges(node)
		for edge := range edges {
			dest := edge.Dest()
			if !visited[dest] {
				q.Enqueue(dest)
				visited[dest] = true
			}
		}
	}
}

// DFS returns a depth-first search iteration channel from a node
func (G *Graph) DFS(source int) (chan int, error) {
	if !G.validNode(source) {
		return nil, errors.New("The source is not a valid node in the graph")
	}
	ch := make(chan int, 20)
	visited := make([]bool, G.Nodes())
	go func(source int) {
		G.dfsHelper(source, ch, visited)
		for i, value := range visited {
			if !value {
				G.dfsHelper(i, ch, visited)
			}
		}
		close(ch)
	}(source)
	return ch, nil
}

func (G *Graph) dfsHelper(source int, ch chan int, visited []bool) {
	stack := linear.NewStack()
	log.Println("fresh Pushing ", source)
	stack.Push(source)
	for stack.Size() != 0 {
		cur, _ := stack.Pop()
		node := cur.(int)
		if !visited[node] {
			ch <- node
			log.Println("DFS EDGE: ", node)
			visited[node] = true
		}
		edges, _ := G.AdjEdges(node)
		for edge := range edges {
			dest := edge.Dest()
			if !visited[dest] {
				log.Println("Pushing ", dest)
				stack.Push(dest)
			}
		}
	}
}
