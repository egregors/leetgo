package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validateStackSequences(t *testing.T) {
	assert.True(t, validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	assert.True(t, validateStackSequences([]int{2, 1, 0}, []int{1, 2, 0}))
	assert.False(t, validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
