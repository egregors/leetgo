package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isCompleteTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				NewTreeNode("[1,2,3,4,5,6]"),
			},
			true,
		},
		{
			"2",
			args{
				NewTreeNode("[1,2,3,4,5,null,7]"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isCompleteTree(tt.args.root), "isCompleteTree(%v)", tt.args.root)
		})
	}
}
