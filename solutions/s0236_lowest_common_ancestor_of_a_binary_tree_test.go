package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor236(t *testing.T) {
	tree := &TreeNode{
		3,
		&TreeNode{
			5,
			&TreeNode{
				6,
				nil,
				nil,
			},
			&TreeNode{
				2,
				&TreeNode{
					7,
					nil,
					nil,
				},
				&TreeNode{
					4,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			1,
			&TreeNode{
				0,
				nil,
				nil,
			},
			&TreeNode{
				8,
				nil,
				nil,
			},
		},
	}
	assert.Equal(t, 3, lowestCommonAncestor236(tree, tree.Left, tree.Right).Val)
}
