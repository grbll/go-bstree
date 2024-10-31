package avltree

import (
	"errors"
	// "fmt"

	ibs "github.com/grbll/go-bstree/internal/internalbstree"
)

type avlData[B ibs.Comparable[B]] struct {
	value   *B
	balance *int
}

func newAVLData[B ibs.Comparable[B]](value B) *avlData[B] {
	return &avlData[B]{value: &value, balance: func(val int) *int { return &val }(0)}
}

func (data avlData[T]) Value() *T {
	return data.value
}

func (data avlData[T]) Balance() *int {
	return data.balance
}

type avlNode[B ibs.Comparable[B]] struct {
	ibs.TreeNode[avlData[B], B]
}

func NewAVLNode[B ibs.Comparable[B]](value B) *avlNode[B] {
	return &avlNode[B]{*ibs.NewTreeNode(*newAVLData(value))}
}

func NewAVLTree[B ibs.Comparable[B]](values []B) (*avlNode[B], error) {
	if len(values) != 0 {
		tree := NewAVLNode(values[0])

		for _, value := range values[1:] {
			tree.Insert(*newAVLData(value))
		}

		return tree, nil
	} else {
		return nil, errors.New("Needs at leas one value to create Tree.")
	}
}
