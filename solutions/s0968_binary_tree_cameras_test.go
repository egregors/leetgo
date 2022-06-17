package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCameraCover(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 0",
			args{root: NewTreeNode("[0,0,null,0,0]")},
			1,
		},
		{
			"Test 1",
			args{root: NewTreeNode("[0,0,null,0,null,0,null,null,0]")},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minCameraCover(tt.args.root), "minCameraCover(%v)", tt.args.root)
		})
	}
}
