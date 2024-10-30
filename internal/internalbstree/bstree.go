package bstree

import (
	"fmt"
	"strings"
)

// BSTrees are filled with Comparable values
type Comparable[T any] interface {
	Less(T) bool
}

// base struct for holding values
type TreeNode[T Comparable[T]] struct {
	value  *T
	parent *TreeNode[T]
	left   *TreeNode[T]
	right  *TreeNode[T]
}

// tree to string method
func (tree *TreeNode[T]) String() string {

	if tree != nil {
		sb := strings.Builder{}
		sb.WriteString(fmt.Sprintf("%v", *tree.value))
		sb.WriteByte('[')
		sb.WriteString(tree.left.String())
		sb.WriteByte(',')
		sb.WriteString(tree.right.String())
		sb.WriteByte(']')

		return sb.String()
	} else {
		return fmt.Sprintf("%v", nil)
	}
}

// constructor for new nodes
func NewTreeNode[T Comparable[T]]() *TreeNode[T] {
	return &TreeNode[T]{}
}

// insert method based on comparison
func (tree *TreeNode[T]) Insert(value T) {
	if tree.value != nil {

		var head **TreeNode[T] = &tree

		for {
			if value.Less(*(*head).value) { //a smaller value should go to the right side
				if (*head).left != nil {
					head = &(*head).left //if left child is not nil, we keep looking...
				} else {
					(*head).left = &TreeNode[T]{value: &value, parent: *head, left: nil, right: nil} //here we can place our new node using the correct parent
					return
				}
			} else {
				if (*head).right != nil { //...and a larger one to the right
					head = &(*head).right //if right child is not nil, we keep looking...
				} else {
					(*head).right = &TreeNode[T]{value: &value, parent: *head, left: nil, right: nil} //here we can place our new node using the correct parent
					return
				}
			}
		}

	} else {
		tree.value = &value // this should only happen for the root of a newly initiated tree
	}
	return

}

type BSTree[T Comparable[T]] interface {
	Insert(T)
}

func NewBSTree[T Comparable[T]](values []T, create func() *TreeNode[T]) *TreeNode[T] {
	tree := create()

	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}
