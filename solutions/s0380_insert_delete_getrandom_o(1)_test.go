package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomizedSet(t *testing.T) {
	rs := NewRandomizedSet()
	assert.True(t, rs.Insert(1))
	assert.False(t, rs.Remove(2))
	assert.True(t, rs.Insert(2))
	assert.Contains(t, []int{1, 2}, rs.GetRandom())
	assert.True(t, rs.Remove(1))
	assert.False(t, rs.Insert(2))
	assert.Equal(t, 2, rs.GetRandom())
}
