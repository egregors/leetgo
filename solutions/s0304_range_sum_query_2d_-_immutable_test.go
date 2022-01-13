package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase(t *testing.T) {
	m := NewNumMatrix(
		[][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		},
	)
	assert.Equal(t, 8, m.SumRegion(2, 1, 4, 3))
	assert.Equal(t, 11, m.SumRegion(1, 1, 2, 2))
	assert.Equal(t, 12, m.SumRegion(1, 2, 2, 4))

	m = NewNumMatrix(
		[][]int{
			{-4, -5},
		},
	)
	assert.Equal(t, -4, m.SumRegion(0, 0, 0, 0))
	assert.Equal(t, -9, m.SumRegion(0, 0, 0, 1))
	assert.Equal(t, -5, m.SumRegion(0, 1, 0, 1))
}
