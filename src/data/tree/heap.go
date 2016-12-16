package tree

// Heap represents a min-heap backed by tree
type Heap struct {
	tree  Tree
	count int
}

// NewHeap return a new Heap
func NewHeap() *Heap {
	heap := new(Heap)
	return heap
}

// Root return the root or the min element
func (H *Heap) Root() *Node {
	return H.tree.Root()
}

// Count returns the number of elements in the Heap
func (H *Heap) Count() int {
	return H.count
}

// Tree returns the backing tree of the heap
func (H *Heap) Tree() Tree {
	return H.tree
}

// SetRoot sets the root of the tree
func (H *Heap) SetRoot(node *Node) {
	H.tree.SetRoot(node)
}

// PreOrderIterator returns  a channel of *Node which is pre-order traversal
func (H *Heap) PreOrderIterator() chan *Node {
	return H.tree.PreOrderIterator()
}

// InOrderIterator returns  a channel of *Node which is in-order traversal
func (H *Heap) InOrderIterator() chan *Node {
	return H.tree.InOrderIterator()
}

// PostOrderIterator returns a channel of *Node which is post-order traversal
func (H *Heap) PostOrderIterator() chan *Node {
	return H.tree.PostOrderIterator()
}
