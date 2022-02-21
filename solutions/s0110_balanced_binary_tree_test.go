package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 1",
			args{root: NewTreeNode("[3,9,20,null,null,15,7]")},
			true,
		},
		{
			"Test 2",
			// fixme: https://github.com/egregors/TreeNode/issues/1
			args{root: NewTreeNode("[1,2,2,3,3,null,null,4,4]")},
			false,
		},
		{
			"Test 3",
			args{root: nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isBalanced(tt.args.root), "isBalanced(%v)", tt.args.root)
		})
	}
}
