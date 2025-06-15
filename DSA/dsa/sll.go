package dsa

import (
	"errors"
	"fmt"
)

type Sll[T comparable] struct {
	head   *SllNode[T]
	tail   *SllNode[T]
	length uint
}

func (sll *SllNode[T]) CreateSll(headData T) *Sll[T] {
	newSll := &Sll[T]{head: nil, tail: nil, length: 0}

	if &headData != nil {
		newSll.head = &SllNode[T]{data: headData, next: nil}
	}

	return newSll
}

func (sll *Sll[T]) Append(data T) {
	newSllNode := SllNode[T]{data: data, next: nil}
	if sll.head == nil {
		sll.head = &newSllNode
		sll.tail = &newSllNode
	}

	sll.tail.next = &newSllNode
	sll.length += 1
}

func (sll *Sll[T]) Insert(pos uint, data T) error {
	if pos < 0 {
		return errors.New(fmt.Sprintf("[pos = %d] pos < 0\n", pos))
	}

	if pos > sll.length {
		return errors.New(fmt.Sprintf("[pos - %d ; length - %d] : pos > length\n", pos, sll.length))
	}

	newSllNode := &SllNode[T]{data: data, next: nil}

	if pos == 0 {
		newSllNode.next = sll.head
		if sll.head.next == nil {
			sll.tail = sll.head
		}
		sll.head = newSllNode
		sll.length += 1
		return nil
	}

	if pos == sll.length {
		sll.tail.next = newSllNode
		sll.tail = newSllNode
		sll.length += 1
		return nil
	}

	start := sll.head
	var index uint = 0
	for index+1 < pos {
		start = start.next
		index += 1
	}

	newSllNode.next = start.next
	start.next = newSllNode
	sll.length += 1
	return nil
}

func (sll *Sll[T]) DeleteAt(pos uint) error {
	if pos >= sll.length {
		return errors.New(fmt.Sprintf("[pos = %d; length = %d] pos >= length\n", pos, sll.length))
	}

	if pos == 0 {
		sll.head = sll.head.next
		if sll.head == nil {
			sll.tail = nil
		}

		sll.length -= 1
		return nil
	}

	start := sll.head
	var index uint = 0
	for index+1 < pos {
		start = start.next
		index += 1
	}

	if pos == sll.length-1 {
		sll.tail = start
		sll.tail.next = nil
		sll.length -= 1
		return nil
	}

	start.next = start.next.next
	sll.length -= 1

	return nil
}

func (sll *Sll[T]) Delete(data T) error {

	if sll.head.data == data {
		sll.head = sll.head.next
		if sll.head == nil {
			sll.tail = nil
		}

		sll.length -= 1
		return nil
	}

	start := sll.head
	var index uint = 0
	for index+1 < sll.length {
		if start.next.data == data {
			if index == (sll.length - 1) {
				sll.tail = start
				sll.tail.next = nil
				sll.length -= 1
				return nil
			} else {
				start.next = start.next.next
				sll.length -= 1
				return nil
			}
		}
		start = start.next
		index += 1
	}

	return errors.New("data not found in the list")
}

func (sll *Sll[T]) Search(data T) (uint, error) {
	start := sll.head
	var index uint = 0
	for index < sll.length {
		if start.data == data {
			return index, nil
		}

		start = start.next
		index += 1
	}

	return 0, errors.New("data not found in the list")
}

func (sll *Sll[T]) Reverse() {
	if sll.length <= 1 {
		return
	}

	var prev *SllNode[T] = nil
	curr, next := sll.head, sll.head.next

	for curr != nil {
		curr.next = prev
		prev = curr
		curr = next
		if next.next == nil {
			next = next.next
		}
	}

	sll.tail = sll.head
	sll.head = prev
}
