package internalbstree

import (
	"fmt"
	// "strings"
)

type Comparable[T any] interface {
	Less(T) bool
}

type Printable[T any] interface {
	fmt.Stringer
}

type treeNode[T any] interface {
	get() T
	set(T)
	partent() *treeNode[T]
	left() *treeNode[T]
	right() *treeNode[T]
}

type avlNode[T Comparable[T]] interface {
	treeNode[T]
	balance()
}

type bsTree[T any] interface {
	treeNode[T]
	Insert(T)
}

type bstTreePrintable[T Printable[T]] interface {
	bsTree[T]
	fmt.Stringer
}

func NewBSTree[T any](values []T, create func() bsTree[T]) bsTree[T] {
	tree := create()

	for _, value := range values {
		tree.Insert(value)
	}

	return tree
}
