package solutions

import (
	"reflect"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"Test 1",
			args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
				val: 5,
			},
			&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
