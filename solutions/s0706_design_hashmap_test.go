package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMyHashMap(t *testing.T) {
	obj := NewMyHashMap()
	obj.Put(1, 1)
	assert.Equal(t, 1, obj.Get(1))

	obj.Put(2, 2)
	assert.Equal(t, 1, obj.Get(1))
	assert.Equal(t, 2, obj.Get(2))
	assert.Equal(t, -1, obj.Get(3))

	obj.Put(2, 1)
	assert.Equal(t, 1, obj.Get(2))

	obj.Remove(2)
	assert.Equal(t, -1, obj.Get(2))
}
