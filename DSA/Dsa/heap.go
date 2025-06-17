package dsa

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	arr    []T
	Length uint
	Min    bool
}

func (h *Heap[T]) CreateHeap(initialSize uint, isMinHeap bool) *Heap[T] {
	_h := Heap[T]{arr: make([]T, initialSize), Length: 0, Min: isMinHeap}
	return &_h
}

func (h *Heap[T]) Insert(data T) {
	h.arr = append(h.arr, data)
	h.Length += 1
	_ = h.bubbleUp(h.Length - 1)
}

func (h *Heap[T]) Delete(data T) {
	index, err := h.Search(data)
	if err == nil {
		panic("data not present in the list")
	}

	h.swap(index, h.Length-1)
	index = h.bubbleUp(index)
	h.bubbleDown(index)
}

func (h *Heap[T]) DeleteRoot() T {
	res := h.arr[0]
	h.swap(0, h.Length-1)
	h.arr = h.arr[:h.Length-1]
	h.bubbleDown(0)
	return res
}

func (h *Heap[T]) swap(index1 uint, index2 uint) {
	temp := h.arr[index1]
	h.arr[index1] = h.arr[index2]
	h.arr[index2] = temp
}

func (h *Heap[T]) Search(data T) (uint, error) {
	var low uint = 0
	var high uint = h.Length - 1
	var mid uint
	for low <= high {
		mid = low + ((high - low) / 2)
		if h.arr[mid] == data {
			return mid, nil
		} else if h.arr[mid] > data {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, errors.New("data not found in the heap")
}

func (h *Heap[T]) cmp(data1 T, data2 T) bool {
	if h.Min {
		if data1 < data2 {
			return true
		} else {
			return false
		}
	} else {
		if data1 > data2 {
			return true
		} else {
			return false
		}
	}
}

func (h *Heap[T]) bubbleUp(index uint) uint {
	for {
		var parentIndex uint = (index - 1) / 2
		if h.cmp(h.arr[index], h.arr[parentIndex]) {
			h.swap(index, parentIndex)
			index = parentIndex
		} else {
			break
		}
	}

	return index
}

func (h *Heap[T]) bubbleDown(index uint) {

	var leftChild uint
	var rightChild uint
	var res uint

	for {
		leftChild = 2*index + 1
		rightChild = 2*index + 2
		if leftChild > h.Length-1 && rightChild > h.Length-1 {
			break
		}

		if leftChild < h.Length-1 && h.cmp(h.arr[leftChild], h.arr[index]) {
			res = leftChild
		}

		if rightChild < h.Length-1 && h.cmp(h.arr[rightChild], h.arr[index]) {
			res = rightChild
		}

		if res == index {
			break
		}

		h.swap(index, res)
		index = res
	}
}

func (h *Heap[T]) String() string {
	return fmt.Sprintf("%+v", *h)
}
