package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				root: NewTreeNode("[1,2,3,4,5,6]"),
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{root: NewTreeNode("[1]")},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countNodes(tt.args.root), "countNodes(%v)", tt.args.root)
		})
	}
}
