package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMode(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			"case 1",
			args{
				NewTreeNode("[1,null,2,2]"),
			},
			[]int{2},
		},
		{
			"case 2",
			args{
				NewTreeNode("[0,-1,2,-1,1]"),
			},
			[]int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantRes, findMode(tt.args.root), "findMode(%v)", tt.args.root)
		})
	}
}
