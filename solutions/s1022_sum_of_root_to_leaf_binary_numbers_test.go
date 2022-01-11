package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumRootToLeaf(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sumRootToLeaf(tt.args.root), "sumRootToLeaf(%v)", tt.args.root)
		})
	}
}
