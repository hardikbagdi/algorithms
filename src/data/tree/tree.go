package tree

// Tree represents a binary tree data type
type Tree struct {
	root *Node
}

// NewTree returns a new empty Tree
func NewTree() *Tree {
	return new(Tree)
}

// Root returns the root node of the tree
func (T *Tree) Root() *Node {
	return T.root
}

// SetRoot sets the root of the tree
func (T *Tree) SetRoot(node *Node) {
	T.root = node
}

// Contains returns true if the elements is present in the tree
func (T *Tree) Contains(value int) bool {
	for node := range T.PreOrderIterator() {
		if node.Value() == value {
			//FIXME What happens to goroutine and channel;Garbage generated?
			return true
		}
	}
	return false
}

// PreOrderIterator returns  a channel of *Node which is pre-order traversal
func (T *Tree) PreOrderIterator() chan *Node {
	//TODO when changing to general trees; change channel to  interface{}
	ch := make(chan *Node, 10)
	go func() {
		helperPreOrder(T.Root(), ch)
		close(ch)
	}()
	return ch
}

func helperPreOrder(node *Node, ch chan *Node) {
	if node == nil {
		return
	}
	ch <- node
	helperPreOrder(node.Left(), ch)
	helperPreOrder(node.Right(), ch)
}

// InOrderIterator returns  a channel of *Node which is in-order traversal
func (T *Tree) InOrderIterator() chan *Node {
	ch := make(chan *Node, 10)
	go func() {
		helperInOrder(T.Root(), ch)
		close(ch)
	}()
	return ch
}

func helperInOrder(node *Node, ch chan *Node) {
	if node == nil {
		return
	}
	helperInOrder(node.Left(), ch)
	ch <- node
	helperInOrder(node.Right(), ch)
}

// PostOrderIterator returns a channel of *Node which is post-order traversal
func (T *Tree) PostOrderIterator() chan *Node {
	ch := make(chan *Node, 10)
	go func() {
		helperPostOrder(T.Root(), ch)
		close(ch)
	}()
	return ch
}

func helperPostOrder(node *Node, ch chan *Node) {
	if node == nil {
		return
	}
	helperPostOrder(node.Left(), ch)
	helperPostOrder(node.Right(), ch)
	ch <- node
}
