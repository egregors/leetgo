package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDiffInBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				NewTreeNode("[4,2,6,1,3]"),
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minDiffInBST(tt.args.root), "minDiffInBST(%v)", tt.args.root)
		})
	}
}
