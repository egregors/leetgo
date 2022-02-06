package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 1",
			args{
				root:    NewTreeNode("[3,4,5,1,2,null,null]"),
				subRoot: NewTreeNode("[4,1,2]"),
			},
			true,
		},
		{
			"Test 2",
			args{
				root:    NewTreeNode("[3,4,5,1,2,null,null,null,null,0,null]"),
				subRoot: NewTreeNode("[4,1,2]"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isSubtree(tt.args.root, tt.args.subRoot), "isSubtree(%v, %v)", tt.args.root, tt.args.subRoot)
		})
	}
}
