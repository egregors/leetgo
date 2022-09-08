package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pruneTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"Test 1",
			args{root: NewTreeNode("[1,null,0,0,1]")},
			NewTreeNode("[1,null,0,null,1]"),
		},
		{
			"Test 2",
			args{root: NewTreeNode("[1,0,1,0,0,0,1]")},
			NewTreeNode("[1,null,1,null,1]"),
		},
		{
			"Test 3",
			args{root: NewTreeNode("[1,1,0,1,1,0,1,0]")},
			NewTreeNode("[1,1,0,1,1,null,1]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, pruneTree(tt.args.root), "pruneTree(%v)", tt.args.root)
		})
	}
}
