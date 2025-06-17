package dsa

type StackDll[T comparable] struct {
	DllLst *Dll[T]
	Length uint
}

func (s *StackDll[T]) CreateStackDll() *StackDll[T] {
	stk := &StackDll[T]{DllLst: &Dll[T]{Head: nil, Length: 0}}
	return stk
}

func (s *StackDll[T]) Push(data T) {
	s.DllLst.Insert(0, data)
	s.Length += 1
}

func (s *StackDll[T]) Pop() T {
	if s.Length <= 0 {
		panic("nothing in stack to pop")
	}
	res := s.DllLst.Head.Data
	s.DllLst.DeleteAt(0)
	s.Length -= 1
	return res
}

func (s *StackDll[T]) Top() T {
	return s.DllLst.Head.Data
}
