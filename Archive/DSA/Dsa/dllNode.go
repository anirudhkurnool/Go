package dsa

type DllNode[T comparable] struct {
	Data T
	Prev *DllNode[T]
	Next *DllNode[T]
}

func (d *DllNode[T]) CreateDllNode(data T) *DllNode[T] {
	newDllNode := DllNode[T]{Data: data, Next: nil}
	return &newDllNode
}
