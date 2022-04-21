package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMyHashSet(t *testing.T) {
	set := NewMyHashSet()
	set.Add(1)
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(3))
	set.Add(2)
	assert.True(t, set.Contains(2))
	set.Remove(2)
	assert.False(t, set.Contains(2))
}
