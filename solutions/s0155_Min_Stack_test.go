package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	s := NewMinStack()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	assert.Equal(t, -3, s.GetMin())
	s.Pop()
	assert.Equal(t, 0, s.Top())
	assert.Equal(t, -2, s.GetMin())
}
