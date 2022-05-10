package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum3(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 4}}, combinationSum3(3, 7))
	assert.Equal(t, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}, combinationSum3(3, 9))
	assert.Equal(t, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}, combinationSum3(9, 45))
}
