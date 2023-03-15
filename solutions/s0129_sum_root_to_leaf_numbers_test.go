package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumNumbers(t *testing.T) {
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
			args{root: NewTreeNode("[1,2,3]")},
			25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sumNumbers(tt.args.root), "sumNumbers(%v)", tt.args.root)
		})
	}
}
