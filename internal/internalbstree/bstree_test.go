// some tests written by openai
package bstree_test

import (
	"testing"

	ibs "github.com/grbll/go-bstree/internal/internalbstree"
)

// Mock type to implement Comparable interface
type MyInt int

func (d MyInt) Less(value MyInt) bool {
	return d < value
}

// Mock implementation of DataInterface
type IntHolder struct {
	value MyInt
}

// Implement the Value method for IntHolder
func (i IntHolder) Value() *MyInt {
	return &i.value
}

// Test for creating a new tree node
func TestNewTreeNode(t *testing.T) {
	holder := &IntHolder{value: 10}
	node := ibs.NewTreeNode(holder)

	if (*node.Data).Value() == nil || *(*node.Data).Value() != holder.value {
		t.Errorf("Expected node value to be %d, got %v", holder.value, (*node.Data).Value())
	}
}

// Test for inserting values into the tree
func TestTreeNodeInsert(t *testing.T) {
	rootHolder := IntHolder{value: 10}
	root := ibs.NewTreeNode(rootHolder)

	// Insert values into the tree
	valuesToInsert := []MyInt{5, 15, 3, 7, 12, 18, 1, 4, 6, 8} // Total of 10 nodes
	for _, val := range valuesToInsert {
		root.Insert(IntHolder{value: val})
	}

	// Check the structure of the tree
	if root.Left == nil || *(*root.Left.Data).Value() != 5 {
		t.Errorf("Expected left child value to be 5, got %v", (*root.Left.Data).Value())
	}
	if root.Right == nil || *(*root.Right.Data).Value() != 15 {
		t.Errorf("Expected right child value to be 15, got %v", (*root.Right.Data).Value())
	}
	if root.Left.Left == nil || *(*root.Left.Left.Data).Value() != 3 {
		t.Errorf("Expected left-left child value to be 3, got %v", (*root.Left.Left.Data).Value())
	}
	if root.Left.Right == nil || *(*root.Left.Right.Data).Value() != 7 {
		t.Errorf("Expected left-right child value to be 7, got %v", (*root.Left.Right.Data).Value())
	}
	if root.Right.Left == nil || *(*root.Right.Left.Data).Value() != 12 {
		t.Errorf("Expected right-left child value to be 12, got %v", (*root.Right.Left.Data).Value())
	}
	if root.Right.Right == nil || *(*root.Right.Right.Data).Value() != 18 {
		t.Errorf("Expected right-right child value to be 18, got %v", (*root.Right.Right.Data).Value())
	}
	if root.Left.Left.Left == nil || *(*root.Left.Left.Left.Data).Value() != 1 {
		t.Errorf("Expected left-left-left child value to be 1, got %v", (*root.Left.Left.Left.Data).Value())
	}
	if root.Left.Left.Right == nil || *(*root.Left.Left.Right.Data).Value() != 4 {
		t.Errorf("Expected left-left-right child value to be 4, got %v", (*root.Left.Left.Right.Data).Value())
	}
	if root.Left.Right.Left == nil || *(*root.Left.Right.Left.Data).Value() != 6 {
		t.Errorf("Expected left-right-left child value to be 6, got %v", (*root.Left.Right.Left.Data).Value())
	}
	if root.Left.Right.Right == nil || *(*root.Left.Right.Right.Data).Value() != 8 {
		t.Errorf("Expected left-right-right child value to be 8, got %v", (*root.Left.Right.Right.Data).Value())
	}
}

// Test for string representation of the tree
func TestTreeNodeString(t *testing.T) {
	holder := &IntHolder{value: 10}
	root := ibs.NewTreeNode(holder)

	// Insert values into the tree
	valuesToInsert := []MyInt{5, 15, 3, 7} // Total of 4 nodes
	for _, val := range valuesToInsert {
		root.Insert(&IntHolder{value: val})
	}

	expectedString := "10[5[3[<nil>,<nil>],7[<nil>,<nil>]],15[<nil>,<nil>]]"
	if root.String() != expectedString {
		t.Errorf("Expected string representation to be %s, got %s", expectedString, root.String())
	}
}
