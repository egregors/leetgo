package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeatManager(t *testing.T) {
	m := NewSeatManager(5)
	assert.Equal(t, 1, m.Reserve())
	assert.Equal(t, 2, m.Reserve())
	m.Unreserve(2)
	assert.Equal(t, 2, m.Reserve())
	assert.Equal(t, 3, m.Reserve())
	assert.Equal(t, 4, m.Reserve())
	assert.Equal(t, 5, m.Reserve())
	m.Unreserve(5)
}
