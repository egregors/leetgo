package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pathSum(t *testing.T) {
	tree := &TreeNode{
		5,
		&TreeNode{
			4,
			&TreeNode{
				11,
				&TreeNode{
					7,
					nil,
					nil,
				},
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
			nil,
		},
		&TreeNode{
			8,
			&TreeNode{
				13,
				nil,
				nil,
			},
			&TreeNode{
				4,
				&TreeNode{
					5,
					nil,
					nil,
				},
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
		},
	}
	res := pathSum(tree, 22)
	assert.EqualValues(t, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}, res)
}
