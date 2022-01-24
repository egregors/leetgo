package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tree := &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}
	res := zigzagLevelOrder(tree)
	assert.EqualValues(t, [][]int{{3}, {20, 9}, {15, 7}}, res)
}
