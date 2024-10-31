package bstree

import (
	"fmt"
	"strings"
)

type DataInterface[T any] interface {
	Less(T) bool
}

// base struct for holding values
type TreeNode[T DataInterface[T]] struct {
	Data   *T
	Parent *TreeNode[T]
	Left   *TreeNode[T]
	Right  *TreeNode[T]
}

// tree to string method
func (tree *TreeNode[T]) String() string {

	if tree != nil {
		sb := strings.Builder{}
		sb.WriteString(fmt.Sprintf("%v", *tree.Data))
		sb.WriteByte('[')
		sb.WriteString(tree.Left.String())
		sb.WriteByte(',')
		sb.WriteString(tree.Right.String())
		sb.WriteByte(']')

		return sb.String()
	} else {
		return fmt.Sprintf("%v", nil)
	}
}

// constructor for new nodes
func NewTreeNode[T DataInterface[T]]() *TreeNode[T] {
	return &TreeNode[T]{}
}

// insert method based on comparison
func (tree *TreeNode[T]) Insert(value T) {
	if tree.Data != nil {

		var head **TreeNode[T] = &tree

		for {
			if value.Less(*(*head).Data) { //a smaller value should go to the right side
				if (*head).Left != nil {
					head = &(*head).Left //if left child is not nil, we keep looking...
				} else {
					(*head).Left = &TreeNode[T]{Data: &value, Parent: *head, Left: nil, Right: nil} //here we can place our new node using the correct parent
					return
				}
			} else {
				if (*head).Right != nil { //...and a larger one to the right
					head = &(*head).Right //if right child is not nil, we keep looking...
				} else {
					(*head).Right = &TreeNode[T]{Data: &value, Parent: *head, Left: nil, Right: nil} //here we can place our new node using the correct parent
					return
				}
			}
		}

	} else {
		tree.Data = &value // this should only happen for the root of a newly initiated tree
	}
	return

}

type BSTree[T DataInterface[T]] interface {
	Insert(T)
}

func NewBSTree[T DataInterface[T]](values []T, create func() BSTree[T]) BSTree[T] {
	tree := create()

	for _, value := range values {
		tree.Insert(value)
	}

	return tree
}
