package dsa

import "fmt"

type CircularQueue[T comparable] struct {
	arr    []T
	Capcty int
	Length uint
	front  int
	back   int
}

func (cq *CircularQueue[T]) CreateCircularQueue(size int) *CircularQueue[T] {
	if size <= 0 {
		panic("Queue capcity cannot be negitive")
	}
	_cq := CircularQueue[T]{arr: make([]T, size), Capcty: size, Length: 0, front: 0, back: 0}
	return &_cq
}

func (cq *CircularQueue[T]) Enqueue(data T) {
	if (cq.back-cq.front+1 == cq.Capcty) || (cq.back-cq.front == -1) {
		panic("queue is full")
	}

	if cq.front == cq.back {
		//list is empty
		cq.arr[cq.back] = data
		cq.Length += 1
		return
	}

	curr := (cq.back + 1) % cq.Capcty
	cq.back = curr
	cq.arr[curr] = data
	cq.Length += 1
}

func (cq *CircularQueue[T]) Dequeue() T {
	if cq.front == cq.back {
		panic("queue is empty")
	}

	res := cq.arr[cq.front]
	cq.front = (cq.front + 1) % cq.Capcty
	return res
}

func (cq *CircularQueue[T]) Peek() T {
	return cq.arr[cq.front]
}

func (cq *CircularQueue[T]) String() string {
	return fmt.Sprintf("%+v", *cq)
}
