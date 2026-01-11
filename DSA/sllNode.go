package dsa

type SllNode[T any] struct {
	data int
	next *SllNode[T]
}
