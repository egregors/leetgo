package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	assert.Equal(t, 3, tree.Val)
	assert.Equal(t, 9, tree.Left.Val)
	assert.Equal(t, 20, tree.Right.Val)
	assert.Equal(t, 15, tree.Right.Left.Val)
	assert.Equal(t, 7, tree.Right.Right.Val)
}
