// some tests written by openai
package bstree

import (
	"testing"
)

// A sample type that implements the Comparable interface
type MyInt int

func (a MyInt) Less(b MyInt) bool {
	return a < b
}

// Test for inserting a single value
func TestInsertSingleValue(t *testing.T) {
	tree := NewBSTree([]MyInt{5}, NewTreeNode[MyInt])
	if tree.value == nil || *tree.value != 5 {
		t.Errorf("Expected root value to be 5, got %v", tree.value)
	}
}

// Test for inserting multiple values
func TestInsertMultipleValues(t *testing.T) {
	values := []MyInt{5, 3, 7, 2, 4, 6, 8}
	tree := NewBSTree(values, NewTreeNode[MyInt])

	if tree.value == nil || *tree.value != 5 {
		t.Errorf("Expected root value to be 5, got %v", tree.value)
	}
	if tree.left.value == nil || *tree.left.value != 3 {
		t.Errorf("Expected left child to be 3, got %v", tree.left.value)
	}
	if tree.right.value == nil || *tree.right.value != 7 {
		t.Errorf("Expected right child to be 7, got %v", tree.right.value)
	}
}

// Test for String method
func TestStringMethod(t *testing.T) {
	values := []MyInt{5, 3, 7}
	tree := NewBSTree(values, NewTreeNode[MyInt])
	expectedString := "5[3[<nil>,<nil>],7[<nil>,<nil>]]" // adjust according to your string representation logic

	if tree.String() != expectedString {
		t.Errorf("Expected string representation to be %s, got %s", expectedString, tree.String())
	}
}
