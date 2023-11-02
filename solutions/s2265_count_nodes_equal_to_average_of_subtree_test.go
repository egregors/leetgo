package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_averageOfSubtree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				root: NewTreeNode("[4,8,5,0,1,null,6]"),
			},
			5,
		},
		{
			"test 2",
			args{
				root: NewTreeNode("[1]"),
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, averageOfSubtree(tt.args.root), "averageOfSubtree(%v)", tt.args.root)
		})
	}
}
