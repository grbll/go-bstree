package tree

type Node[T any] interface {
	Parent() *T
	Left() *T
	Right() *T
}

func depthFirstTraversel[T Node[T], C any](tree *T, hook func(node *T, direction uint8, outputPointer *C) C) C {
	var out C
	var direction uint8 = 0 //0-> left child next, 1->right child next, 2->parent next
	if tree != nil {
		for (*tree).Parent() != nil && direction != 2 {
			if direction == 0 {
				out = hook(tree, 0, &out)

				if (*tree).Left() != nil {
					tree = (*tree).Left()
				} else {
					out = hook(nil, 0, &out)
					direction = 1
				}

			} else if direction == 1 {
				out = hook(tree, 1, &out)

				if (*tree).Right() != nil {
					tree = (*tree).Right()
					direction = 0
				} else {
					out = hook(nil, 0, &out)
					direction = 2
				}

			} else { //direction == 2
				out = hook(tree, 2, &out)

				if tree == (*(*tree).Parent()).Left() {
					direction = 1
				} else {
					direction = 2
				}

				tree = (*tree).Parent()

			}
		}
		out = hook(tree, 2, &out)
	} else {
		out = hook(nil, direction, &out)
	}
	return out
}

// func depthFirstTraversel[T DataInterface[B], B Comparable[B], C any](tree *TreeNode[T, B], hook func(node *TreeNode[T, B], direction int, outputPointer *C)) C {
// 	var output C
// 	direction := 0
// 	if tree != nil {
// 		for tree.Parent != nil || direction != 2 {
// 			if direction == 0 {
//
// 				hook(tree, 0, &output)
//
// 				if tree.Left != nil {
// 					tree = tree.Left
// 				} else {
// 					hook(nil, 0, &output)
// 					direction = 1
// 				}
//
// 			} else if direction == 1 {
// 				hook(tree, 1, &output)
//
// 				if tree.Right != nil {
// 					tree = tree.Right
// 					direction = 0
// 				} else {
// 					hook(nil, 0, &output)
// 					direction = 2
// 				}
//
// 			} else { //direction == 2
//
// 				hook(tree, 2, &output)
// 				if tree == tree.Parent.Left {
// 					direction = 1
// 				} else {
// 					direction = 2
// 				}
// 				tree = tree.Parent
//
// 			}
// 		}
// 		hook(tree, 2, &output)
// 	} else {
// 		hook(nil, direction, &output)
// 	}
// 	return output
// }
//
// func stringHook[T DataInterface[B], B Comparable[B]](node *TreeNode[T, B], direction int, stringBuilder *strings.Builder) {
// 	if stringBuilder == nil {
// 		stringBuilder = &strings.Builder{}
// 	}
// 	if node != nil {
// 		if direction == 0 {
// 			if node.Data != nil && (*node.Data).Value() != nil {
// 				stringBuilder.WriteString(fmt.Sprintf("%v", *(*node.Data).Value()))
// 			} else {
// 				stringBuilder.WriteString(fmt.Sprintf("%v", nil))
// 			}
// 			stringBuilder.WriteByte('[')
// 		} else if direction == 1 {
// 			stringBuilder.WriteByte(',')
// 		} else { // direction == 2
// 			stringBuilder.WriteByte(']')
// 		}
// 		return
//
// 	} else { //node == nil
// 		return
// 	}
// }
//
// func sliceHook[T DataInterface[B], B Comparable[B]](node *TreeNode[T, B], direction int, slice *[]B) {
// 	if slice == nil {
// 		slice = &[]B{}
// 	}
// 	if node != nil {
// 		if direction == 1 {
// 			if node.Data != nil && (*node.Data).Value() != nil {
// 				*slice = (append(*slice, *(*node.Data).Value()))
// 				return
// 			} else {
// 				return
// 			}
// 		}
// 	} else { //node==nil
// 		return
// 	}
// }
//
// // insert method based on comparison
// func (tree *TreeNode[T, B]) Insert(data T) {
// 	var head **TreeNode[T, B] = &tree
//
// 	for {
// 		if (*head).Data != nil {
// 			if (*data.Value()).Less(*(*(*head).Data).Value()) { //a smaller value should go to the right side
// 				if (*head).Left != nil {
// 					head = &(*head).Left //if left child is not nil, we keep looking...
// 				} else {
// 					(*head).Left = NewTreeNode(data) //here we can place our new node using the correct parent
// 					(*head).Left.Parent = *head
// 					return
// 				}
// 			} else {
// 				if (*head).Right != nil { //...and a larger one to the right
// 					head = &(*head).Right //if right child is not nil, we keep looking...
// 				} else {
// 					(*head).Right = NewTreeNode(data) //here we can place our new node using the correct parent
// 					(*head).Right.Parent = (*head)
// 					return
// 				}
// 			}
// 		} else {
// 			(*head).Data = &data
// 		}
// 	}
// }
//
// // tree to string method
// func (tree *TreeNode[T, B]) String() string {
// 	sb := depthFirstTraversel(tree, stringHook)
// 	return sb.String()
// }
//
// func (tree *TreeNode[T, B]) Slice() []B {
// 	return depthFirstTraversel(tree, sliceHook)
// }
