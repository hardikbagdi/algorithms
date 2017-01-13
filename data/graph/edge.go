package graph

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
