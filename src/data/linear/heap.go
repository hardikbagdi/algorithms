package linear

import "errors"

// Heap represents a backed backed by a slice
type Heap struct {
	array []Comparable
	count int
}

// NewHeap returns a new empty heap
func NewHeap() *Heap {
	heap := new(Heap)
	heap.array = make([]Comparable, 20)
	heap.count = 0
	return heap
}

// Count returns the number of elements currently in the heap
func (H *Heap) Count() int {
	return H.count
}

// Search determines if an element exists in the heap or not
func (H *Heap) Search(value Comparable) bool {

	return false
}

// ElementAt returns element at index. Sets bool if element at index exists; else false
func (H *Heap) ElementAt(index int) (Comparable, bool) {
	return nil, false
}

// Insert adds the value to the heap; expanding the underlying structure if necessary
func (H *Heap) Insert(value Comparable) error {
	if H.Search(value) {
		return errors.New("Value already exists in the heap.")
	}
	//first search for the element; if it exists then return the error
	//cause we don't use the first index
	if H.count+1 == len(H.array) {
		//expand array
		H.resize()
	}
	H.count++
	H.array[H.count] = value
	H.swimUp(H.count - 1)
	return nil
}

// Remove removes the value from the heap
func (H *Heap) Remove(value Comparable) error {
	return nil

}

// GetRoot returns min or max depending upon the implementaion of Compare of the user defined type
func (H *Heap) GetRoot() (Comparable, bool) {
	if H.count == 0 {
		return nil, false
	}
	return H.array[1], true

}

func (H *Heap) resize() {

}

func (H *Heap) swimUp(index int) {

}

func (H *Heap) sinkDown(index int) {

}
