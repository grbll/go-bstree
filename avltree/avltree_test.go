// openai tests
package avltree_test

import (
	"fmt"
	"testing"

	avl "github.com/grbll/go-bstree/avltree"
)

// Mock type to implement the Comparable interface
type MyInt int

func (a MyInt) Less(b MyInt) bool {
	return a < b
}

// Test for creating a new AVL tree with one value
func TestNewAVLTree_SingleValue(t *testing.T) {
	values := []MyInt{10}
	root, err := avl.NewAVLTree(values)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if root == nil {
		t.Fatal("Expected root to be non-nil")
	}
	if *root.Data.Value() != 10 {
		t.Errorf("Expected root value to be 10, got %v", *root.Data.Value())
	}
	if *root.Data.Balance() != 0 {
		t.Errorf("Expected root balance to be 0, got %v", *root.Data.Balance())
	}
}

// Test for creating a new AVL tree with multiple values
func TestMultipleValuesInsert(t *testing.T) {
	values := []MyInt{10, 20, 30, 25, 15}
	root, err := avl.NewAVLTree(values)
	fmt.Printf("%v", root)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if root == nil {
		t.Fatal("Expected root to be non-nil")
	}

	// Check root value
	if *root.Data.Value() != 10 {
		t.Errorf("Expected root value to be 10, got %v", *root.Data.Value())
	}

	// Check the left child and its children
	if root.Right == nil || *root.Right.Data.Value() != 20 {
		t.Errorf("Expected right child value to be 20, got %v", *root.Right.Data.Value())
	}
	if root.Right.Right == nil || *root.Right.Right.Data.Value() != 30 {
		t.Errorf("Expected right-right child value to be 30, got %v", *root.Right.Right.Data.Value())
	}
	if root.Right.Left == nil || *root.Right.Left.Data.Value() != 15 {
		t.Errorf("Expected right-left child value to be 15, got %v", *root.Right.Left.Data.Value())
	}
	if root.Right.Right.Left == nil || *root.Right.Right.Left.Data.Value() != 25 {
		t.Errorf("Expected right-right-left child value to be 25, got %v", *root.Right.Left.Left.Data.Value())
	}
}

// Test for creating an AVL tree with duplicates
// func TestNewAVLTree_Duplicates(t *testing.T) {
// 	values := []MyInt{10, 20, 10, 30}
// 	root, err := avl.NewAVLTree(values)
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}
//
// 	// Check that root value is the first inserted (10)
// 	if *root.Data.Value() != 10 {
// 		t.Errorf("Expected root value to be 10, got %v", *root.Data.Value())
// 	}
// }
//
// // Test for error handling when creating an empty AVL tree
// func TestNewAVLTree_Empty(t *testing.T) {
// 	_, err := avl.NewAVLTree([]MyInt{})
// 	if err == nil {
// 		t.Fatal("Expected an error for empty input, got nil")
// 	}
// }
//
// // Test AVL tree balance after multiple insertions
// func TestNewAVLTree_Balance(t *testing.T) {
// 	values := []MyInt{10, 20, 30, 40, 50}
// 	root, err := avl.NewAVLTree(values)
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}
//
// 	if root == nil {
// 		t.Fatal("Expected root to be non-nil")
// 	}
//
// 	// Assert root value and check the balance
// 	if *root.Data.Value() != 30 {
// 		t.Errorf("Expected root value to be 30, got %v", *root.Data.Value())
// 	}
// 	if *root.Data.Balance() != 0 {
// 		t.Errorf("Expected root balance to be 0, got %v", *root.Data.Balance())
// 	}
// }
//
// // Test for AVL tree structure after insertions
// func TestAVLTree_Structure(t *testing.T) {
// 	values := []MyInt{50, 30, 70, 20, 40, 60, 80}
// 	root, err := avl.NewAVLTree(values)
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}
//
// 	// Check the structure of the tree
// 	if *root.Data.Value() != 50 {
// 		t.Errorf("Expected root value to be 50, got %v", *root.Data.Value())
// 	}
// 	if root.Left == nil || *root.Left.Data.Value() != 30 {
// 		t.Errorf("Expected left child value to be 30, got %v", *root.Left.Data.Value())
// 	}
// 	if root.Right == nil || *root.Right.Data.Value() != 70 {
// 		t.Errorf("Expected right child value to be 70, got %v", *root.Right.Data.Value())
// 	}
// 	if root.Left.Left == nil || *root.Left.Left.Data.Value() != 20 {
// 		t.Errorf("Expected left-left grandchild value to be 20, got %v", *root.Left.Left.Data.Value())
// 	}
// 	if root.Left.Right == nil || *root.Left.Right.Data.Value() != 40 {
// 		t.Errorf("Expected left-right grandchild value to be 40, got %v", *root.Left.Right.Data.Value())
// 	}
// 	if root.Right.Left == nil || *root.Right.Left.Data.Value() != 60 {
// 		t.Errorf("Expected right-left grandchild value to be 60, got %v", *root.Right.Left.Data.Value())
// 	}
// 	if root.Right.Right == nil || *root.Right.Right.Data.Value() != 80 {
// 		t.Errorf("Expected right-right grandchild value to be 80, got %v", *root.Right.Right.Data.Value())
// 	}
// }
