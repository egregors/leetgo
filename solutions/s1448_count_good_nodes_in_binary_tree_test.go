package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_goodNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root: NewTreeNode("[3,1,4,3,null,1,5]"),
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				root: NewTreeNode("[3,3,null,4,2]"),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, goodNodes(tt.args.root), "goodNodes(%v)", tt.args.root)
		})
	}
}
