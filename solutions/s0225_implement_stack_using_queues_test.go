package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyStack(t *testing.T) {
	stack := NewMyStack()
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Top())
	assert.Equal(t, 2, stack.Pop())
	assert.False(t, stack.Empty())
}
