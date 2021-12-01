package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMyHashMap(t *testing.T) {
	obj := NewMyHashMap()
	obj.Put(1, 1)
	assert.Equal(t, []pair{{1, 1}}, obj.xs)
	obj.Put(2, 2)
	assert.Equal(t, []pair{{1, 1}, {2, 2}}, obj.xs)

	assert.Equal(t, 1, obj.Get(1))
	assert.Equal(t, []pair{{1, 1}, {2, 2}}, obj.xs)

	assert.Equal(t, -1, obj.Get(3))
	assert.Equal(t, []pair{{1, 1}, {2, 2}}, obj.xs)

	obj.Put(2, 1)
	assert.Equal(t, []pair{{1, 1}, {2, 1}}, obj.xs)

	assert.Equal(t, 1, obj.Get(2))
	assert.Equal(t, []pair{{1, 1}, {2, 1}}, obj.xs)

	obj.Remove(2)
	assert.Equal(t, []pair{{1, 1}}, obj.xs)

	assert.Equal(t, -1, obj.Get(2))
	assert.Equal(t, []pair{{1, 1}}, obj.xs)
}
