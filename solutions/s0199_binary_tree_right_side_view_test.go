package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rightSideView(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
	}
	res := rightSideView(tree)
	assert.Equal(t, []int{1, 3, 4}, res)
}
