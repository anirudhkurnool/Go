package dsa

type DllNode[T any] struct {
	data T
	prev *DllNode[T]
	next *DllNode[T]
}
