// some tests written by openai
package bstree_test

import (
	"fmt"
	"testing"

	ibs "github.com/grbll/go-bstree/internal/internalbstree"
)

type IntWrapper struct {
	Value int
}

func (a IntWrapper) String() string {
	return fmt.Sprintf("%v", a.Value)
}

func (a IntWrapper) Less(b IntWrapper) bool {
	return a.Value < b.Value
}
func TestTreeNodeInsert(t *testing.T) {
	root := ibs.NewTreeNode[IntWrapper]()

	values := []IntWrapper{{10}, {5}, {15}, {3}, {7}, {12}, {18}}
	for _, v := range values {
		root.Insert(v)
	}

	// Root node test
	if root.Data == nil || (*root.Data).Value != 10 {
		t.Errorf("Expected root node with value 10, got %v", root.Data)
	}

	// Left and right child tests
	if root.Left == nil || (*root.Left.Data).Value != 5 {
		t.Errorf("Expected left child of root to be 5, got %v", root.Left)
	}
	if root.Right == nil || (*root.Right.Data).Value != 15 {
		t.Errorf("Expected right child of root to be 15, got %v", root.Right)
	}
}

func TestTreeNodeString(t *testing.T) {
	root := ibs.NewTreeNode[IntWrapper]()
	values := []IntWrapper{{10}, {5}, {15}, {3}, {7}, {12}, {18}}
	for _, v := range values {
		root.Insert(v)
	}

	// Expected string output check
	expected := "10[5[3[<nil>,<nil>],7[<nil>,<nil>]],15[12[<nil>,<nil>],18[<nil>,<nil>]]]"
	result := root.String()
	if result != expected {
		t.Errorf("Expected string representation %s, got %s", expected, result)
	}
}
