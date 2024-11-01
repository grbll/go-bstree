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
	Value() *B
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
			if (*data.Value()).Less(*(*(*head).Data).Value()) { //a smaller value should go to the right side
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
		origin := 0
		sb := strings.Builder{}
		head := &tree

		for (*head).Parent != nil || origin != 2 {
			if origin == 0 {
				if (*head).Data != nil && (*(*head).Data).Value() != nil {
					sb.WriteString(fmt.Sprintf("%v", *(*(*head).Data).Value()))
				} else {
					sb.WriteString(fmt.Sprintf("%v", nil))
				}
				sb.WriteByte('[')
				if (*head).Left == nil {
					sb.WriteString(fmt.Sprintf("%v", nil))
					origin = 1
				} else {
					head = &(*head).Left
					origin = 0
				}
			} else if origin == 1 {
				sb.WriteByte(',')
				if (*head).Right == nil {
					sb.WriteString(fmt.Sprintf("%v", nil))
					origin = 2
				} else {
					head = &(*head).Right
					origin = 0
				}
			} else {
				sb.WriteByte(']')
				if (*head).Parent.Left == (*head) {
					origin = 1
				} else {
					origin = 2
				}
				head = &(*head).Parent
			}

		}
		sb.WriteByte(']')

		return sb.String()

		// sb.WriteString(fmt.Sprintf("%v", *(*tree.Data).Value()))
		// sb.WriteByte('[')
		// sb.WriteString(tree.Left.String())
		// sb.WriteByte(',')
		// sb.WriteString(tree.Right.String())
		// sb.WriteByte(']')

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
