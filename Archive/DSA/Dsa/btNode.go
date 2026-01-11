package dsa

import "golang.org/x/exp/constraints"

type BtNode[T constraints.Ordered] struct {
	Data       T
	LeftChild  *BtNode[T]
	RightChild *BtNode[T]
}

func CreateBtNode[T constraints.Ordered](data T) *BtNode[T] {
	return &BtNode[T]{Data: data, LeftChild: nil, RightChild: nil}
}
