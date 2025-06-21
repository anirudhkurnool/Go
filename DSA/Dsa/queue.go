package dsa

type Queue[T comparable] struct {
	Arr    []T
	Length uint
}

func CreateQueue[T comparable](initialSize uint) *Queue[T] {
	_q := &Queue[T]{Arr: make([]T, initialSize), Length: 0}
	return _q
}

func (q *Queue[T]) Append(data T) {
	q.Arr = append(q.Arr, data)
	q.Length += 1
}

func (q *Queue[T]) PopLeft() T {
	if q.Length <= 0 {
		panic("no data in queue to pop")
	}
	res := q.Arr[0]
	q.Arr = q.Arr[1:]
	q.Length -= 1
	return res
}

func (q *Queue[T]) Peek() T {
	if q.Length <= 0 {
		panic("no data in queue to pop")
	}
	return q.Arr[0]
}
