package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumArray(t *testing.T) {
	obj := NewNumArray([]int{1, 3, 5})
	assert.Equal(t, 9, obj.SumRange(0, 2))
	obj.Update(1, 2)
	assert.Equal(t, 8, obj.SumRange(0, 2))
}
