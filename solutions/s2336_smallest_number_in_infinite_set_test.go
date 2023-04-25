package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSmallestInfiniteSet(t *testing.T) {
	s := NewSmallestInfiniteSet()
	s.AddBack(2)
	assert.Equal(t, 1, s.PopSmallest())
	assert.Equal(t, 2, s.PopSmallest())
	assert.Equal(t, 3, s.PopSmallest())
	s.AddBack(1)
	assert.Equal(t, 1, s.PopSmallest())
	assert.Equal(t, 4, s.PopSmallest())
	assert.Equal(t, 5, s.PopSmallest())
}
