package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root  *TreeNode
		val   int
		depth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 0",
			args{root: NewTreeNode("[4,2,6,3,1,5]"), val: 1, depth: 2},
			NewTreeNode("[4,1,1,2,null,null,6,3,1,5]").String(),
		},
		{
			"Test 1",
			args{root: NewTreeNode("[3,1,null,null,2]"), val: 1, depth: 3},
			NewTreeNode("[3,1,null,1,1,null,null,null,2]").String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				addOneRow(tt.args.root,
					tt.args.val,
					tt.args.depth).String(),
				"addOneRow(%v, %v, %v)", tt.args.root, tt.args.val, tt.args.depth)
		})
	}
}
