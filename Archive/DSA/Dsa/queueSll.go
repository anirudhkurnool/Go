package dsa

type QueueSll[T comparable] struct {
	SllQ   *Sll[T]
	Length uint
}

func (q *QueueSll[T]) CreateQueueSll() *QueueSll[T] {
	_q := QueueSll[T]{SllQ: &Sll[T]{Head: nil, Length: 0}, Length: 0}
	return &_q
}

func (q *QueueSll[T]) Append(data T) {
	q.SllQ.Append(data)
	q.Length += 1
}

func (q *QueueSll[T]) PopLeft() T {
	res := q.SllQ.Head.Data
	q.SllQ.DeleteAt(0)
	q.Length -= 1
	return res
}

func (q *QueueSll[T]) Peek() T {
	res := q.SllQ.Head.Data
	return res
}
