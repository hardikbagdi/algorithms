package linear

import (
	"io/ioutil"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func verifyHeap(heap *Heap) bool {
	for i := 1; i < heap.count; i++ {
		l := heap.left(i)
		r := heap.right(i)
		if l < heap.count {
			if heap.array[i] < heap.array[l] {
				continue
			} else {
				log.Println("Violation:", heap.array[i], "not less than", heap.array[l])
				return false
			}
		}
		if r < heap.count {
			if heap.array[i] < heap.array[r] {
				continue
			} else {
				log.Println("Violation:", heap.array[i], "not less than", heap.array[r])
				return false
			}
		}
	}
	return true
}

func TestHeap(T *testing.T) {
	log.SetOutput(ioutil.Discard)
	var ele int
	log.Println("Heap test enter")
	t := time.Now()
	rand.Seed(t.Unix())
	heap := NewHeap()
	assert.Equal(T, 0, heap.Count())

	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(8)
	heap.Insert(2)
	ret := verifyHeap(heap)
	assert.Equal(T, true, ret)
	ele, _ = heap.ElementAt(0)
	assert.Equal(T, 2, ele)
	log.Println(heap.Array())

	index, ok := heap.Search(8)
	assert.Equal(T, 2, index)
	assert.Equal(T, true, ok)

	log.Println("decrease key called")
	heap.DecreaseKey(8, 1)
	log.Println(heap)

	ele, _ = heap.ElementAt(0)
	assert.Equal(T, 1, ele)

	err := heap.DeleteRoot()
	assert.Nil(T, err)

	ele, _ = heap.ElementAt(0)
	assert.Equal(T, 2, ele)

	index, ok = heap.Search(8)
	assert.Equal(T, 0, index)
	assert.Equal(T, false, ok)

	index, ok = heap.Search(10)
	assert.Equal(T, true, ok)
	assert.Equal(T, 2, index)

	ele, ok = heap.ElementAt(10)
	assert.Equal(T, false, ok)
	assert.Equal(T, 0, ele)

	err = heap.RemoveAt(1)
	assert.Nil(T, err)
	log.Println(heap)
	ret = verifyHeap(heap)
	assert.Equal(T, true, ret)

	for i := 100; i < 120; i++ {
		heap.Insert(i)
	}
	assert.Equal(T, 22, heap.Count())
	log.Println(heap)
	ret = verifyHeap(heap)
	assert.Equal(T, true, ret)

	heap.Remove(100)
	heap.RemoveAt(1000)
	heap.RemoveAt(4)
	heap.DeleteRoot()
	ret = verifyHeap(heap)
	assert.Equal(T, true, ret)

	log.Println(heap)
	heap.IncreaseKey(10, 10000)
	ret = verifyHeap(heap)
	assert.Equal(T, true, ret)

	heap.DecreaseKey(116, 1)
	log.Println(heap)
	ret = verifyHeap(heap)
	assert.Equal(T, true, ret)

	for j := 1; j < 10000; j++ {
		heap2 := NewHeap()
		heap = heap2
		for i := 1; i < 1000; i++ {
			//			log.Println("Inserting")
			heap.Insert(rand.Intn(1000))
			ret = verifyHeap(heap)
			assert.Equal(T, true, ret)
		}
		for i := 1; i < 1000; i++ {
			//			log.Println("Removing")
			//			log.Println(heap)
			heap.Remove(rand.Intn(1000))
			ret = verifyHeap(heap)
			if !ret {
				log.Println("==========================")
				//				log.Println(heap)
				return
			}
			assert.Equal(T, true, ret)
		}
	}
	log.Println(heap)

	log.Println("Heap test finished")
}

func TestNaiveMergeHeap(T *testing.T) {
	var ret bool
	heap := NewHeap()
	heap2 := NewHeap()
	for i := 1; i < 10; i++ {
		heap.Insert(i)
	}
	for i := 10; i < 20; i++ {
		heap2.Insert(i)
	}
	ret = verifyHeap(heap)
	assert.Equal(T, true, ret)
	ret = verifyHeap(heap2)
	assert.Equal(T, true, ret)

	heap.NaiveMergeHeap(heap2)

	ret = verifyHeap(heap)
	assert.Equal(T, true, ret)
}

func TestMergeHeap(T *testing.T) {
	var ret bool
	heap := NewHeap()
	heap2 := NewHeap()
	for i := 1; i < 10; i++ {
		heap.Insert(i)
	}
	for i := 10; i < 20; i++ {
		heap2.Insert(i)
	}
	ret = verifyHeap(heap)
	assert.Equal(T, true, ret)
	ret = verifyHeap(heap2)
	assert.Equal(T, true, ret)

	heap.MergeHeap(heap2)
	heap.EnsureCapacity(100)

	ret = verifyHeap(heap)
	assert.Equal(T, true, ret)
}
