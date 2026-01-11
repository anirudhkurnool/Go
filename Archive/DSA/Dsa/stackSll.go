package dsa

type StackSll[T comparable] struct {
	SllLst *Sll[T]
	Length uint
}

func (s *StackSll[T]) CreateStackSll() *StackSll[T] {
	stk := &StackSll[T]{SllLst: &Sll[T]{Head: nil, Length: 0}}
	return stk
}

func (s *StackSll[T]) Push(data T) {
	s.SllLst.Insert(0, data)
	s.Length += 1
}

func (s *StackSll[T]) Pop() T {
	if s.Length <= 0 {
		panic("nothing in stack to pop")
	}
	res := s.SllLst.Head.Data
	s.SllLst.DeleteAt(0)
	s.Length -= 1
	return res
}

func (s *StackSll[T]) Top() T {
	return s.SllLst.Head.Data
}
