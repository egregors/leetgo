package solutions

import "testing"

func Test_findTarget(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test",
			args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   2,
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
						Val:  6,
						Left: nil,
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
				k: 9,
			},
			true,
		},
		{
			"Test",
			args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   2,
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
						Val:  6,
						Left: nil,
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
				k: 28,
			},
			false,
		},
		{
			"Test",
			args{
				root: &TreeNode{
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
				k: 4,
			},
			true,
		},
		{
			"Test",
			args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   -5,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
				},
				k: 10,
			},
			true,
		},
		{
			"Test",
			args{
				root: &TreeNode{
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
				k: 1,
			},
			false,
		},
		{
			"Test",
			args{
				root: &TreeNode{
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
				k: 3,
			},
			true,
		},
		{
			"Test",
			args{
				root: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				k: 2,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
