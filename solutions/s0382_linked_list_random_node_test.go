package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution_GetRandom(t *testing.T) {
	tree := NewListNode([]int{1, 2, 3})
	solution := NewSolution(tree)
	assert.NotEqual(t, []int{1, 2, 3}, solution.GetRandom())
}
