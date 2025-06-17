package dsa

import (
	"errors"
	"fmt"
)

type Dll[T comparable] struct {
	Head   *DllNode[T]
	Tail   *DllNode[T]
	Length uint
}

func (dll *Dll[T]) CreateDll() *Dll[T] {
	dll1 := &Dll[T]{Head: nil, Tail: nil, Length: 0}
	return dll1
}

func (dll *Dll[T]) Append(data T) {
	newNode := &DllNode[T]{Data: data, Prev: nil, Next: nil}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		dll.Length += 1
		return
	}

	dll.Tail.Next = newNode
	newNode.Prev = dll.Tail
	dll.Tail = dll.Tail.Next
	dll.Length += 1
}

func (dll *Dll[T]) Insert(pos uint, data T) {
	if pos > dll.Length {
		panic(fmt.Sprintf("[(pos = %d) > (length = %d)]", pos, dll.Length))
	}

	newNode := &DllNode[T]{Data: data, Prev: nil, Next: nil}

	if pos == 0 {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
		dll.Length += 1
		return
	}

	var index uint = 0
	start := dll.Head

	for index+1 < pos {
		start = start.Next
		index += 1
	}

	newNode.Next = start.Next
	newNode.Prev = start
	start.Next = newNode
	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	} else {
		dll.Tail = newNode
	}
	dll.Length += 1
}

func (dll *Dll[T]) DeleteAt(pos uint) {
	if pos >= dll.Length {
		panic(fmt.Sprintf("[(pos = %d) > (length = %d)]", pos, dll.Length))
	}

	if pos == 0 {
		dll.Head = dll.Head.Next
		if dll.Head == nil {
			dll.Tail = nil
		} else {
			dll.Head.Prev = nil
		}

		dll.Length -= 1
		return
	}

	var index uint = 0
	start := dll.Head

	for index+1 < pos {
		index += 1
		start = start.Next
	}

	start.Next = start.Next.Next
	start.Next.Prev = nil
	start.Next.Next = nil
	if start.Next != nil {
		start.Next.Prev = start
	} else {
		dll.Tail = start
	}

	dll.Length -= 1
}

func (dll *Dll[T]) Search(data T) (uint, error) {
	var index uint = 0
	start := dll.Head
	for index < dll.Length {
		if start.Data == data {
			return index, nil
		}
		index += 1
		start = start.Next
	}

	return 0, errors.New("data not found")
}

func (dll *Dll[T]) Delete(data T) {
	index, error := dll.Search(data)
	if error != nil {
		panic("data not present in the list")
	}

	dll.DeleteAt(index)
}

func (dll *Dll[T]) Reverse() {
	if dll.Length <= 1 {
		return
	}

	var prev *DllNode[T] = nil
	curr, next := dll.Head, dll.Head.Next

	for curr != nil {
		curr.Next = prev
		prev = curr
		curr = next
		if next != nil {
			next = next.Next
		}
	}

	dll.Tail = dll.Head
	dll.Head = prev
}

func (dll *Dll[T]) ToSlice() []T {
	slice := make([]T, dll.Length)
	start := dll.Head
	for start != nil {
		slice = append(slice, start.Data)
	}

	return slice
}
