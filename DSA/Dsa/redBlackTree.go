package dsa

import (
	"math"

	"golang.org/x/exp/constraints"
)

type color int

const (
	RED color = iota
	BLACK
)

type RbtNode[T constraints.Ordered] struct {
	Data       T
	Clr        color
	LeftChild  *RbtNode[T]
	RightChild *RbtNode[T]
}

func CreateRbtNode[T constraints.Ordered](data T) *RbtNode[T] {
	return &RbtNode[T]{Data: data, Clr: BLACK, LeftChild: nil, RightChild: nil}
}

type Rbt[T constraints.Ordered] struct {
	Root     *RbtNode[T]
	NumNodes uint
}

func (rbt *Rbt[T]) CreateRbt(rootData T) *Rbt[T] {
	return &Rbt[T]{Root: &RbtNode[T]{Data: rootData, Clr: RED, LeftChild: nil, RightChild: nil}, NumNodes: 1}
}

func (rbt *Rbt[T]) InsertNodeRbt(data T) {
	rbt.Root = InsertRecursiveRbt[T](rbt.Root, data)
	rbt.NumNodes += 1
	rbt.Root = rbt.insertBalance(rbt.Root)
}

func (rbt *Rbt[T]) DeleteNodeRbt(data T) {
	var deletedNodeColor color
	rbt.Root = DeleteRecursiveRbt[T](rbt.Root, data, &deletedNodeColor)
	rbt.NumNodes -= 1
	if deletedNodeColor == BLACK {
		rbt.Root = rbt.deleteBalance(rbt.Root)
	}
}

func (rbt *Rbt[T]) leftRotate(parent *RbtNode[T]) {
	res := parent.RightChild
	parent.RightChild = res.LeftChild
	res.LeftChild = parent
}

func (rbt *Rbt[T]) rightRotate(parent *RbtNode[T]) {
	res := parent.LeftChild
	parent.LeftChild = res.RightChild
	res.RightChild = parent
}

func (rbt *Rbt[T]) balanceFactor(node *RbtNode[T]) int {
	return int(HeightRbt(node.LeftChild)) - int(HeightRbt(node.RightChild))
}

func InsertRecursiveRbt[T constraints.Ordered](root *RbtNode[T], data T) *RbtNode[T] {
	if root == nil {
		return CreateRbtNode[T](data)
	}

	if root.Data > data {
		root.LeftChild = InsertRecursiveRbt(root.LeftChild, data)
	} else if root.Data < data {
		root.RightChild = InsertRecursiveRbt(root.RightChild, data)
	}

	return root
}

func (avl *Rbt[T]) Delete(data T) {
	avl.Root = DeleteRecursiveRbt[T](avl.Root, data)
	avl.NumNodes -= 1
}

func DeleteRecursiveRbt[T constraints.Ordered](root *RbtNode[T], data T) *RbtNode[T] {
	if root == nil {
		return nil
	}

	if root.Data == data {
		if root.LeftChild == nil && root.RightChild == nil {
			return nil
		} else if root.LeftChild != nil && root.RightChild != nil {
			root.Data = LeftMostNodeRbt(root.RightChild).Data
			root.RightChild = DeleteRecursiveRbt[T](root.RightChild, root.Data)
		} else if root.LeftChild != nil {
			return root.LeftChild
		} else if root.RightChild != nil {
			return root.RightChild
		}
	}

	if root.Data > data {
		root.LeftChild = DeleteRecursiveRbt[T](root.LeftChild, data)
	} else if root.Data < data {
		root.RightChild = DeleteRecursiveRbt[T](root.RightChild, data)
	}

	return root
}

func LeftMostNodeRbt[T constraints.Ordered](start *RbtNode[T]) *RbtNode[T] {
	if start == nil {
		return nil
	}
	for start.LeftChild != nil {
		start = start.LeftChild
	}

	return start
}

func HeightRbt[T constraints.Ordered](start *RbtNode[T]) uint {
	if start == nil {
		return 0
	}

	return uint(math.Max(float64(HeightRbt(start.LeftChild)), float64(HeightRbt(start.RightChild))))
}
