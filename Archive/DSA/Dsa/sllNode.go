package dsa

type SllNode[T comparable] struct {
	Data T
	Next *SllNode[T]
}
