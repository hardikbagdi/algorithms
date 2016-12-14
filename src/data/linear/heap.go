package linear

import (
	"bytes"
	"errors"
	"strconv"
)

// Heap represents a min heap backed backed by a slice
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

func (H *Heap) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Count: ")
	buffer.WriteString(strconv.Itoa(H.count) + "\n[")
	for i := 1; i <= H.count; i++ {
		buffer.WriteString(strconv.Itoa(H.array[i]) + ", ")
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Count returns the number of elements currently in the heap
func (H *Heap) Count() int {
	return H.count
}

// Array return the heap
func (H *Heap) Array() []int {
	return H.array[1 : H.count+1]
}

// Search determines if an element exists in the heap or not; index return(0-based)
func (H *Heap) Search(value int) (int, bool) {
	//TODO improve; use tree pre-order and cut-off
	for i := 1; i <= H.count; i++ {
		if H.array[i] == value {
			return i - 1, true
		}
	}
	return 0, false
}

// ElementAt returns element at index(0-based). Sets bool if element at index exists; else false
func (H *Heap) ElementAt(index int) (int, bool) {
	if index >= H.count {
		return 0, false
	}
	return H.array[index+1], true
	return 0, false
}

// Insert adds the value to the heap; expanding the underlying structure if necessary
func (H *Heap) Insert(value int) error {
	//first search for the element; if it exists then return the error
	_, present := H.Search(value)
	if present {
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

// RemoveAt removes the element at index specified. Index is 0-based
func (H *Heap) RemoveAt(index int) error {
	index++
	if index > H.count {
		return errors.New("Index out of boud")
	}
	H.swap(H.count, index)
	H.array[H.count] = 0
	H.count--
	H.minHeapify(index)
	if H.count < len(H.array)/4 {
		H.resize()
	}
	return nil
}

//EnsureCapacity ensures that the backing array has atleast length of capacity
func (H *Heap) EnsureCapacity(capacity int) {
	if len(H.array) < capacity {
		var newArray []int
		newArray = make([]int, capacity)
		copy(newArray, H.array)
		H.array = newArray
	}
}

// DeleteRoot remove the root elements; throws error if heap is empty
func (H *Heap) DeleteRoot() error {
	if H.count < 1 {
		return errors.New("No element in heap")
	}
	return H.RemoveAt(0)
}

// Remove removes the value from the heap
func (H *Heap) Remove(value int) error {
	index, present := H.Search(value)
	if !present {
		return errors.New("Value not present in heap")
	}
	return H.RemoveAt(index)
}

// GetRoot returns min element
func (H *Heap) GetRoot() (int, bool) {
	if H.count == 0 {
		return 0, false
	}
	//cause we don't use the first index
	return H.array[1], true
}

// DecreaseKey reduces the key to a new value
func (H *Heap) DecreaseKey(oldValue, newValue int) error {
	if newValue >= oldValue {
		return errors.New("New value cannot be greater than or equal to old value")
	}
	index, present := H.Search(oldValue)
	if !present {
		return errors.New("Value not present in heap")
	}
	index++
	H.array[index] = newValue
	H.swimUp(index)
	return nil
}

// IncreaseKey increases  the value of a key
func (H *Heap) IncreaseKey(oldValue, newValue int) error {
	if newValue <= oldValue {
		return errors.New("New value cannot be less than or equal to old value")
	}
	index, present := H.Search(oldValue)
	if !present {
		return errors.New("Value not present in heap")
	}
	H.array[index] = newValue
	H.sinkDown(index)
	return nil
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

func (H *Heap) left(index int) int {
	return 2 * index
}
func (H *Heap) right(index int) int {
	return 2*index + 1
}
func (H *Heap) minHeapify(index int) {
	l := H.left(index)
	r := H.right(index)
	min := index
	if l < H.count+1 && H.array[l] < H.array[min] {
		min = l
	}
	if r < H.count+1 && H.array[r] < H.array[min] {
		min = r
	}
	if index != min {
		H.swap(min, index)
		H.minHeapify(min)
	}
}

func (H *Heap) sinkDown(index int) {
	H.minHeapify(index)
}

func (H *Heap) NaiveMergeHeap(heap1 *Heap) {
	//some checking
	if heap1 == nil {
		return
	}
	array := heap1.Array()
	for value := range array {
		H.Insert(value)
	}
}

func (H *Heap) MergeHeap(heap1 *Heap) {
	//some checking
	if heap1 == nil {
		return
	}
	array1 := heap1.Array()

	H.EnsureCapacity(H.Count() + heap1.Count())
	index := H.Count() + 1
	for i := 0; i < heap1.Count(); i++ {
		H.array[index] = array1[i]
		index++
	}
	H.count = H.Count() + heap1.Count()

	for i := H.count / 2; i > 0; i-- {
		H.minHeapify(i)
	}
}
