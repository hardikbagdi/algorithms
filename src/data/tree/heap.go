package tree

import (
	"errors"
	"strconv"
)

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

// Contains returns true if the elements is present in the tree
func (H *Heap) Contains(value int) bool {
	for node := range H.PreOrderIterator() {
		if node.Value() == value {
			return true
		}
	}
	return false
}

// NodeWithValue does the same thing as search but return the pointer to the node
//TODO decide if this should really be exported or contains could return bool and node pointer
func (H *Heap) NodeWithValue(value int) *Node {
	for node := range H.PreOrderIterator() {
		if node.Value() == value {
			return node
		}
	}
	return nil
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

// Insert adds the value to the heap
func (H *Heap) Insert(value int) error {
	present := H.Contains(value)
	if present {
		return errors.New("Value already present in heap")
	}
	//insert at the last node
	node := new(Node)
	node.SetValue(value)
	H.appendLast(node)
	H.count++
	//swimUp
	H.swimUp(node)
	return nil
}

// DeleteRoot removes the root elements; throws error if heap is empty
func (H *Heap) DeleteRoot() error {
	if H.count < 1 {
		return errors.New("No element in heap")
	}
	H.Remove(H.Root().Value())
	return nil
}

// Remove removes the value from the heap
func (H *Heap) Remove(value int) error {
	node := H.NodeWithValue(value)
	if node == nil {
		return errors.New("Value not present in heap")
	}
	//TODO
	//TODO test for only one node and with two or there nodes
	//get last node
	lastNode := H.getLastNode()
	//swap with last node
	H.swap(lastNode, node)
	//remove the last node from the tree
	parentLastNode := lastNode.Parent()
	// only one node is present in the tree
	if parentLastNode == nil {
		H.SetRoot(nil)
		H.count--
		return nil
	}
	if parentLastNode.Left() == lastNode {
		parentLastNode.SetLeft(nil)
	} else {
		parentLastNode.SetRight(nil)
	}
	lastNode.SetParent(nil)
	H.count--
	//sink down the 'node'
	H.sinkDown(node)
	return nil
}

// DecreaseKey reduces the key to a new value
func (H *Heap) DecreaseKey(oldValue, newValue int) error {
	if newValue >= oldValue {
		return errors.New("New value cannot be greater than or equal to old value")
	}
	node := H.NodeWithValue(oldValue)
	if node == nil {
		return errors.New("Value not present in heap")
	}
	node.SetValue(newValue)
	H.swimUp(node)
	return nil
}

// IncreaseKey increases  the value of a key
func (H *Heap) IncreaseKey(oldValue, newValue int) error {
	if newValue <= oldValue {
		return errors.New("New value cannot be less than or equal to old value")
	}
	node := H.NodeWithValue(oldValue)
	if node == nil {
		return errors.New("Value not present in heap")
	}
	node.SetValue(newValue)
	H.sinkDown(node)
	return nil
}

// swap will swap the values inside the nodes and not the nodes
func (H *Heap) swap(a, b *Node) {
	if a == nil || b == nil {
		return
	}
	backup := a.Value()
	a.SetValue(b.Value())
	b.SetValue(backup)
}

func (H *Heap) swimUp(node *Node) {
	if node == nil {
		return
	}
	if node.Parent() == nil {
		return
	}
	if node.Parent().Value() > node.Value() {
		H.swap(node, node.Parent())
		H.swimUp(node.Parent())
	}
}

func (H *Heap) sinkDown(node *Node) {
	if node == nil {
		return
	}
	min := node
	if node.Left() != nil && node.Left().Value() < min.Value() {
		min = node.Left()
	}
	if node.Right() != nil && node.Right().Value() < min.Value() {
		min = node.Right()
	}
	if min != node {
		H.swap(min, node)
		H.sinkDown(min)
	}
}

// TODO cache the last node rather than caculating every time
func (H *Heap) appendLast(node *Node) {
	c := H.count
	c++
	if c == 1 {
		H.SetRoot(node)
		return
	}
	cur := H.Root()
	if c > 3 {
		binary := intToBinaryString(c)
		for i := 1; i < len(binary)-1; i++ {
			if binary[i] == '0' {
				cur = cur.Left()
			} else {
				cur = cur.Right()
			}
		}
	}
	if c%2 == 0 {
		cur.SetLeft(node)
	} else {
		cur.SetRight(node)
	}
	node.SetParent(cur)
}

func (H *Heap) getLastNode() *Node {
	//TODO get ride of this special case
	//if H.Count == 1 {
	//	return H.Root()
	//}
	binary := intToBinaryString(H.Count())
	cur := H.Root()
	for i := 1; i < len(binary); i++ {
		if binary[i] == '0' {
			cur = cur.Left()
		} else {
			cur = cur.Right()
		}
	}
	return cur
}
func intToBinaryString(number int) string {
	str := ""
	index := 0
	for number > 0 {
		str = strconv.Itoa(number%2) + str
		index++
		number /= 2
	}
	return str
}
