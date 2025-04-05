package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lcaDeepestLeaves(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want string
	}{
		{
			name: "Test 1",
			root: NewTreeNode("[3,5,1,6,2,0,8,null,null,7,4]"),
			want: NewTreeNode("[2,7,4]").String(),
		},
		{
			name: "Test 2",
			root: NewTreeNode("[1]"),
			want: NewTreeNode("[1]").String(),
		},
		{
			name: "Test 3",
			root: NewTreeNode("[0,1,3,null,2]"),
			want: NewTreeNode("[2]").String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lcaDeepestLeaves(tt.root).String(), "lcaDeepestLeaves(%v)", tt.root)
		})
	}
}
