package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(xs, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, xs)

	xs = []int{-1, -100, 3, 99}
	rotate(xs, 2)
	assert.Equal(t, []int{3, 99, -1, -100}, xs)

	xs = []int{-1}
	rotate(xs, 2)
	assert.Equal(t, []int{-1}, xs)

	xs = []int{1, 2}
	rotate(xs, 3)
	assert.Equal(t, []int{2, 1}, xs)
}
