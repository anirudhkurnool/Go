package dsa

type Dll[T comparable] struct {
	head   *DllNode[T]
	tail   *DllNode[T]
	length uint
}

func (d *Dll[T]) CreateDll(headDat T) {
	//
}
