package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	xs := []int{0, 1, 0, 3, 12}
	moveZeroes(xs)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, xs)

	xs = []int{0}
	moveZeroes(xs)
	assert.Equal(t, []int{0}, xs)

}
