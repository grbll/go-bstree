package avltree

import ibs "github.com/grbll/go-bstree/internal/internalbstree"

type avlData[T ibs.DataInterface[T]] struct {
	value   *T
	balance int
}

func (data1 avlData[T]) Less(data2 avlData[T]) bool {
	return (*data1.value).Less(*data2.value)
}

type avlNode[T ibs.DataInterface[T]] struct {
	ibs.TreeNode[avlData[T]]
}

func createAVLNode[T ibs.DataInterface[T]]() *avlNode[T] {
	return &avlNode[T]{ibs.TreeNode[avlData[T]]{Data: &avlData[T]{value: nil, balance: 0}}}
}
