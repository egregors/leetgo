package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthSmallest(t *testing.T) {
	tree := &TreeNode{
		3,
		&TreeNode{
			1,
			nil,
			&TreeNode{
				2,
				nil,
				nil,
			},
		},
		&TreeNode{
			4,
			nil,
			nil,
		},
	}
	assert.Equal(t, 1, kthSmallest(tree, 1))
}
