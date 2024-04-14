package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				NewTreeNode("[3,9,20,null,null,15,7]"),
			},
			24,
		},
		{
			"case 2",
			args{
				NewTreeNode("[1]"),
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sumOfLeftLeaves(tt.args.root), "sumOfLeftLeaves(%v)", tt.args.root)
		})
	}
}
