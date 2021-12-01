package solutions

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Test 1",
			args: args{
				root1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				root2: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  1,
						Left: nil,
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  5,
					Left: nil,
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
