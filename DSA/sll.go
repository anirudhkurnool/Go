package dsa

type SLL[T any] struct {
	head *SllNode[T]
	tail *SllNode[T]
	size int
}

func (sll *SLL[T]) append(newElement T) {
	newNode := &SllNode[T]{data: newElement, next: nil}
	if sll.head == nil {
		sll.head = newNode
		sll.tail = sll.head
	} else {
		sll.tail.next = newNode
		sll.tail = sll.tail.next
	}

	sll.size++
}
