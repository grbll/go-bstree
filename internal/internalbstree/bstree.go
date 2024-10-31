package bstree

import (
	"fmt"
	"strings"
)

// Comparable interface for comparing values
type Comparable[T any] interface {
	Less(T) bool
}

// DataInterface for holding values
type DataInterface[B any] interface {
	Get() *B
	Set(value B)
}

// TreeNode structure holding data
type TreeNode[T DataInterface[B], B Comparable[B]] struct {
	Data   *T
	Parent *TreeNode[T, B]
	Left   *TreeNode[T, B]
	Right  *TreeNode[T, B]
}

// constructor for new nodes
func NewTreeNode[T DataInterface[B], B Comparable[B]](data T) *TreeNode[T, B] {
	return &TreeNode[T, B]{Data: &data, Parent: nil, Left: nil, Right: nil}
}

// insert method based on comparison
func (tree *TreeNode[T, B]) Insert(data T) {
	var head **TreeNode[T, B] = &tree

	for {
		if (*head).Data != nil {
			if (*data.Get()).Less(*(*(*head).Data).Get()) { //a smaller value should go to the right side
				if (*head).Left != nil {
					head = &(*head).Left //if left child is not nil, we keep looking...
				} else {
					(*head).Left = NewTreeNode(data) //here we can place our new node using the correct parent
					(*head).Left.Parent = *head
					return
				}
			} else {
				if (*head).Right != nil { //...and a larger one to the right
					head = &(*head).Right //if right child is not nil, we keep looking...
				} else {
					(*head).Right = NewTreeNode(data) //here we can place our new node using the correct parent
					(*head).Right.Parent = (*head)
					return
				}
			}
		} else {
			(*head).Data = &data
		}
	}
}

// tree to string method
func (tree *TreeNode[T, B]) String() string {

	if tree != nil {
		sb := strings.Builder{}
		sb.WriteString(fmt.Sprintf("%v", *(*tree.Data).Get()))
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

//
// type BSTree[T DataInterface[T]] interface {
// 	Insert(T)
// }
//
// func NewBSTree[T DataInterface[T]](values []T, create func() BSTree[T]) BSTree[T] {
// 	tree := create()
//
// 	for _, value := range values {
// 		tree.Insert(value)
// 	}
//
// 	return tree
// }
