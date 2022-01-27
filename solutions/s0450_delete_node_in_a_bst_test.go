package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	tree := &TreeNode{
		5,
		&TreeNode{
			3,
			&TreeNode{
				2,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
		&TreeNode{
			6,
			nil,
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}
	res := deleteNode(tree, 3)
	assert.Equal(t, 5, res.Val)
	assert.Equal(t, 4, res.Left.Val)
	assert.Equal(t, 6, res.Right.Val)
	assert.Equal(t, 2, res.Left.Left.Val)
	assert.Equal(t, 7, res.Right.Right.Val)
}
