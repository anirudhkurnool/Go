package dsa

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type PriorityQueue[T constraints.Ordered] struct {
	heap   *Heap[T]
	Length uint
}

func (pq *PriorityQueue[T]) CreatePriorityQueue() *PriorityQueue[T] {
	_pq := PriorityQueue[T]{heap: pq.heap.CreateHeap(3, true), Length: 0}
	return &_pq
}

func (pq *PriorityQueue[T]) Insert(data T) {
	pq.heap.Insert(data)
}

func (pq *PriorityQueue[T]) Deque() T {
	return pq.heap.DeleteRoot()
}

func (pq *PriorityQueue[T]) String() string {
	return fmt.Sprintf("%+v", *pq)
}
