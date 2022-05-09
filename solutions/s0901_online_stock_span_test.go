package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockSpanner(t *testing.T) {
	s := NewStockSpanner()
	assert.Equal(t, 1, s.Next(100))
	assert.Equal(t, 1, s.Next(80))
	assert.Equal(t, 1, s.Next(60))
	assert.Equal(t, 2, s.Next(70))
	assert.Equal(t, 1, s.Next(60))
	assert.Equal(t, 4, s.Next(75))
}
