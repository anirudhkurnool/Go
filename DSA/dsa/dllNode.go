package dsa

type DllNode[T comparable] struct {
	data T
	prev *DllNode[T]
	next *DllNode[T]
}

func (d *DllNode[T]) CreateDllNode(data T) *DllNode[T] {
	newDllNode := DllNode[T]{data: data, next: nil}
	return &newDllNode
}
