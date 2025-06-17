package dsa

type QueueDll[T comparable] struct {
	DllQ   *Dll[T]
	Length uint
}

func (q *QueueDll[T]) CreateQueueDll() *QueueDll[T] {
	_q := QueueDll[T]{DllQ: &Dll[T]{Head: nil, Length: 0}, Length: 0}
	return &_q
}

func (q *QueueDll[T]) Append(data T) {
	q.DllQ.Append(data)
	q.Length += 1
}

func (q *QueueDll[T]) PopLeft() T {
	res := q.DllQ.Head.Data
	q.DllQ.DeleteAt(0)
	q.Length -= 1
	return res
}

func (q *QueueDll[T]) Peek() T {
	res := q.DllQ.Head.Data
	return res
}
