package dsa

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Bst[T constraints.Ordered] struct {
	Root     *BtNode[T]
	NumNodes uint
}

func (bst *Bst[T]) CreateBst(data T) *Bst[T] {
	return &Bst[T]{Root: CreateBtNode[T](data), NumNodes: 1}
}

func (bst *Bst[T]) Insert(data T) {
	bst.Root = InsertRecursive[T](bst.Root, data)
	bst.NumNodes += 1
}

func InsertRecursive[T constraints.Ordered](root *BtNode[T], data T) *BtNode[T] {
	if root == nil {
		return CreateBtNode[T](data)
	}

	if root.Data > data {
		root.LeftChild = InsertRecursive(root.LeftChild, data)
	} else if root.Data < data {
		root.RightChild = InsertRecursive(root.RightChild, data)
	}

	return root
}

func (bst *Bst[T]) Delete(data T) {
	bst.Root = DeleteRecursive[T](bst.Root, data)
	bst.NumNodes -= 1
}

func DeleteRecursive[T constraints.Ordered](root *BtNode[T], data T) *BtNode[T] {
	if root == nil {
		return nil
	}

	if root.Data == data {
		if root.LeftChild == nil && root.RightChild == nil {
			return nil
		} else if root.LeftChild != nil && root.RightChild != nil {
			root.Data = LeftMostNode(root.RightChild).Data
			root.RightChild = DeleteRecursive[T](root.RightChild, root.Data)
		} else if root.LeftChild != nil {
			return root.LeftChild
		} else if root.RightChild != nil {
			return root.RightChild
		}
	}

	if root.Data > data {
		root.LeftChild = DeleteRecursive[T](root.LeftChild, data)
	} else if root.Data < data {
		root.RightChild = DeleteRecursive[T](root.RightChild, data)
	}

	return root
}

func LeftMostNode[T constraints.Ordered](start *BtNode[T]) *BtNode[T] {
	if start == nil {
		return nil
	}
	for start.LeftChild != nil {
		start = start.LeftChild
	}

	return start
}

func Height[T constraints.Ordered](start *BtNode[T]) uint {
	if start == nil {
		return 0
	}

	return uint(math.Max(float64(Height(start.LeftChild)), float64(Height(start.RightChild))))
}
