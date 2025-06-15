package dsa

type SllNode[T comparable] struct {
	data T
	next *SllNode[T]
}
