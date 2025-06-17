package dsa

type Stack[T comparable] struct {
	Arr    []T
	Length uint //is this useful???
}

func (s *Stack[T]) CreateStack(initialSize uint) *Stack[T] {
	stack := Stack[T]{Arr: make([]T, initialSize), Length: 0}
	return &stack
}

func (s *Stack[T]) Push(data T) {
	s.Arr = append(s.Arr, data)
	s.Length += 1
}

func (s *Stack[T]) Pop() T {
	if s.Length <= 0 {
		panic("nothing in stack to pop")
	}
	res := s.Arr[s.Length-1]
	s.Arr = s.Arr[:s.Length-1]
	s.Length -= 1
	return res
}

func (s *Stack[T]) Top() T {
	if s.Length <= 0 {
		panic("stack is empty")
	}

	return s.Arr[0]
}
