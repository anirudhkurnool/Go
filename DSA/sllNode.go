package dsa

type SllNode[T any] struct {
	data T
	next *SllNode[T]
}
