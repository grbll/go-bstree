package avltree

import ibs "github.com/grbll/go-bstree/internal/internalbstree"

type Comparable[T any] interface {
	Less(T) bool
}

type avlNode[T Comparable[T]] struct {
	ibs.TreeNode[T]
	balance int
}

func createAVLNode[T Comparable[T]]() *avlNode[T] {
	return &avlNode[T]{*ibs.NewTreeNode[T](), 0}
}
