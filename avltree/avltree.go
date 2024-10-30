package avltree

import ibs "github.com/grbll/go-bstree/internal/internalbstree"

type avlNode[T ibs.Comparable[T]] struct {
	ibs.TreeNode[T]
	balance int
}

func createAVLNode[T ibs.Comparable[T]]() *avlNode[T] {
	return &avlNode[T]{TreeNode: *ibs.NewTreeNode[T](), balance: 0}
}

func (tree *avlNode[T]) Insert(value T) {
	if tree.Value != nil {

		var head **avlNode[T] = &tree

		for {
			if value.Less(*(*head).Value) { //a smaller value should go to the right side
				if (*head).Left != nil {
					head = &(*head).Left //if left child is not nil, we keep looking...
				} else {
					(*head).Left = &TreeNode[T]{Value: &value, Parent: *head, Left: nil, Right: nil} //here we can place our new node using the correct parent
					return
				}
			} else {
				if (*head).Right != nil { //...and a larger one to the right
					head = &(*head).Right //if right child is not nil, we keep looking...
				} else {
					(*head).Right = &TreeNode[T]{Value: &value, Parent: *head, Left: nil, Right: nil} //here we can place our new node using the correct parent
					return
				}
			}
		}

	} else {
		tree.Value = &value // this should only happen for the root of a newly initiated tree
	}
	return

}
func NewAVLTree[T ibs.Comparable[T]](values []T) ibs.BSTree[T] {
	// Use NewBSTree to initialize the tree with the creation function
	tree := ibs.NewBSTree(values, func() ibs.BSTree[T] { return createAVLNode[T]() })
	// Cast the returned TreeNode to *avlNode
	return tree
}
