package linear

import "errors"

//import "strconv"

// Heap represents a backed backed by a slice
type Heap struct {
	array []int
	count int
}

// NewHeap returns a new empty heap
func NewHeap() *Heap {
	heap := new(Heap)
	heap.array = make([]int, 20)
	heap.count = 0
	return heap
}

//func (H *Heap) String() string {
//	s := "Count: "
//	s += strconv.Itoa(H.count) + "\n"
// TODO add array to print
//	return s
//}

// Count returns the number of elements currently in the heap
func (H *Heap) Count() int {
	return H.count
}

// Array return the heap
func (H *Heap) Array() []int {
	return H.array[1 : H.count+1]
}

// Search determines if an element exists in the heap or not
func (H *Heap) Search(value int) bool {
	//TODO improve; use tree pre-order and cut-off
	for i := 1; i <= H.count; i++ {
		if H.array[i] == value {
			return true
		}
	}
	return false
}

// ElementAt returns element at index(0-based). Sets bool if element at index exists; else false
func (H *Heap) ElementAt(index int) (int, bool) {
	if index >= H.count {
		return 0, false
	}
	return H.array[index-1], true
	return 0, false
}

// Insert adds the value to the heap; expanding the underlying structure if necessary
func (H *Heap) Insert(value int) error {
	//first search for the element; if it exists then return the error
	if H.Search(value) {
		return errors.New("Value already exists in the heap.")
	}
	//cause we don't use the first index
	if H.count+1 == len(H.array) {
		//expand array
		H.resize()
	}
	H.count++
	H.array[H.count] = value
	H.swimUp(H.count)
	return nil
}

// Remove removes the value from the heap
func (H *Heap) Remove(value int) error {
	return nil
}

// GetRoot returns min or max depending upon the implementaion of Compare of the user defined type
func (H *Heap) GetRoot() (int, bool) {
	if H.count == 0 {
		return 0, false
	}
	//cause we don't use the first index
	return H.array[1], true
}

func (H *Heap) resize() {
	var newArray []int
	if H.count+1 == len(H.array) {
		//double it
		newArray = make([]int, len(H.array)*2)

	} else if H.count < len(H.array)/4 {
		//reduce to half
		newArray = make([]int, len(H.array)/2)
	}
	//TODO do error checking
	copy(newArray, H.array)
	H.array = newArray
}

func (H *Heap) swap(a, b int) {
	H.array[a], H.array[b] = H.array[b], H.array[a]
}
func (H *Heap) parent(index int) int {
	return index / 2
}

func (H *Heap) swimUp(index int) {
	for index != 1 && H.array[index]-H.array[H.parent(index)] < 0 {
		H.swap(index, H.parent(index))
		index = H.parent(index)
	}
}

func (H *Heap) sinkDown(index int) {
}
