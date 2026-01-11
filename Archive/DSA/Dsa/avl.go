package dsa

import "golang.org/x/exp/constraints"

type Avl[T constraints.Ordered] struct {
	Root     *BtNode[T]
	NumNodes uint
}

func (avl *Avl[T]) CreateAvl(data T) *Avl[T] {
	return &Avl[T]{Root: CreateBtNode[T](data), NumNodes: 1}
}

func (avl *Avl[T]) InsertNodeAvl(data T) {
	avl.Root = InsertRecursive[T](avl.Root, data)
	avl.NumNodes += 1
	avl.Root = avl.balance(avl.Root)
}

func (avl *Avl[T]) DeleteNodeAvl(data T) {
	avl.Root = DeleteRecursive[T](avl.Root, data)
	avl.NumNodes -= 1
	avl.Root = avl.balance(avl.Root)
}

func (avl *Avl[T]) leftRotate(parent *BtNode[T]) {
	res := parent.RightChild
	parent.RightChild = res.LeftChild
	res.LeftChild = parent
}

func (avl *Avl[T]) rightRotate(parent *BtNode[T]) {
	res := parent.LeftChild
	parent.LeftChild = res.RightChild
	res.RightChild = parent
}

func (avl *Avl[T]) balanceFactor(node *BtNode[T]) int {
	return int(Height(node.LeftChild)) - int(Height(node.RightChild))
}

func (avl *Avl[T]) balance(node *BtNode[T]) *BtNode[T] {
	bf := avl.balanceFactor(node)
	if bf < 0 {
		//right heavy
		child_bf := avl.balanceFactor(node.RightChild)
		if child_bf > 0 {
			//right - left case
			avl.rightRotate(node.RightChild)
		}

		avl.leftRotate(node)
	} else {
		//left heavy
		child_bf := avl.balanceFactor(node.LeftChild)
		if child_bf < 0 {
			//left -right case
			avl.leftRotate(node.LeftChild)
		}

		avl.rightRotate(node)
	}

	node.LeftChild = avl.balance(node.LeftChild)
	node.RightChild = avl.balance(node.RightChild)
	return node
}
