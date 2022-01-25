package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBSTIterator(t *testing.T) {
	root := &TreeNode{
		7,
		&TreeNode{
			3,
			nil,
			nil,
		},
		&TreeNode{
			15,
			&TreeNode{
				9,
				nil,
				nil,
			},
			&TreeNode{
				20,
				nil,
				nil,
			},
		},
	}
	i := NewBSTIterator(root)
	assert.Equal(t, 3, i.Next())
	assert.Equal(t, 7, i.Next())
	assert.True(t, i.HasNext())
	assert.Equal(t, 9, i.Next())
	assert.True(t, i.HasNext())
	assert.Equal(t, 15, i.Next())
	assert.True(t, i.HasNext())
	assert.Equal(t, 20, i.Next())
	assert.False(t, i.HasNext())

}
