package dsa

import (
	"errors"
	"fmt"
)

type Sll[T comparable] struct {
	Head   *SllNode[T]
	Tail   *SllNode[T]
	Length uint
}

func (sll *SllNode[T]) CreateSll(HeadData T) *Sll[T] {
	newSll := &Sll[T]{Head: nil, Tail: nil, Length: 0}

	if &HeadData != nil {
		newSll.Head = &SllNode[T]{Data: HeadData, Next: nil}
	}

	return newSll
}

func (sll *Sll[T]) Append(data T) {
	newSllNode := SllNode[T]{Data: data, Next: nil}
	if sll.Head == nil {
		sll.Head = &newSllNode
		sll.Tail = &newSllNode
	}

	sll.Tail.Next = &newSllNode
	sll.Length += 1
}

func (sll *Sll[T]) Insert(pos uint, data T) error {
	if pos < 0 {
		return errors.New(fmt.Sprintf("[pos = %d] pos < 0\n", pos))
	}

	if pos > sll.Length {
		return errors.New(fmt.Sprintf("[pos - %d ; Length - %d] : pos > Length\n", pos, sll.Length))
	}

	newSllNode := &SllNode[T]{Data: data, Next: nil}

	if pos == 0 {
		newSllNode.Next = sll.Head
		if sll.Head.Next == nil {
			sll.Tail = sll.Head
		}
		sll.Head = newSllNode
		sll.Length += 1
		return nil
	}

	if pos == sll.Length {
		sll.Tail.Next = newSllNode
		sll.Tail = newSllNode
		sll.Length += 1
		return nil
	}

	start := sll.Head
	var index uint = 0
	for index+1 < pos {
		start = start.Next
		index += 1
	}

	newSllNode.Next = start.Next
	start.Next = newSllNode
	sll.Length += 1
	return nil
}

func (sll *Sll[T]) DeleteAt(pos uint) error {
	if pos >= sll.Length {
		return errors.New(fmt.Sprintf("[pos = %d; Length = %d] pos >= Length\n", pos, sll.Length))
	}

	if pos == 0 {
		sll.Head = sll.Head.Next
		if sll.Head == nil {
			sll.Tail = nil
		}

		sll.Length -= 1
		return nil
	}

	start := sll.Head
	var index uint = 0
	for index+1 < pos {
		start = start.Next
		index += 1
	}

	if pos == sll.Length-1 {
		sll.Tail = start
		sll.Tail.Next = nil
		sll.Length -= 1
		return nil
	}

	start.Next = start.Next.Next
	sll.Length -= 1

	return nil
}

func (sll *Sll[T]) Delete(data T) error {

	if sll.Head.Data == data {
		sll.Head = sll.Head.Next
		if sll.Head == nil {
			sll.Tail = nil
		}

		sll.Length -= 1
		return nil
	}

	start := sll.Head
	var index uint = 0
	for index+1 < sll.Length {
		if start.Next.Data == data {
			if index == (sll.Length - 1) {
				sll.Tail = start
				sll.Tail.Next = nil
				sll.Length -= 1
				return nil
			} else {
				start.Next = start.Next.Next
				sll.Length -= 1
				return nil
			}
		}
		start = start.Next
		index += 1
	}

	return errors.New("data not found in the list")
}

func (sll *Sll[T]) Search(data T) (uint, error) {
	start := sll.Head
	var index uint = 0
	for index < sll.Length {
		if start.Data == data {
			return index, nil
		}

		start = start.Next
		index += 1
	}

	return 0, errors.New("data not found in the list")
}

func (sll *Sll[T]) Reverse() {
	if sll.Length <= 1 {
		return
	}

	var prev *SllNode[T] = nil
	curr, next := sll.Head, sll.Head.Next

	for curr != nil {
		curr.Next = prev
		prev = curr
		curr = next
		if next.Next == nil {
			next = next.Next
		}
	}

	sll.Tail = sll.Head
	sll.Head = prev
}
