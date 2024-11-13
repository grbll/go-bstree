package tree_test

import (
	"github.com/grbll/go-bstree/tree/main/internal/tree"
	"reflect"
	"testing"
)

func Test_depthFirstTraversel[T any, C any](t *testing.T) {
	type args struct {
		tree *T
		hook func(node *T, direction uint8, outputPointer *C) C
	}
	tests := []struct {
		name string
		args args
		want C
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree.DepthFirstTraversel(tt.args.tree, tt.args.hook); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("depthFirstTraversel() = %v, want %v", got, tt.want)
			}
		})
	}
}
