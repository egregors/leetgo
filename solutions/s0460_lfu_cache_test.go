package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLFUCache(t *testing.T) {
	cache := NewLFUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, 1, cache.Get(1))
	cache.Put(3, 3)
	assert.Equal(t, -1, cache.Get(2))
	assert.Equal(t, 3, cache.Get(3))
	cache.Put(4, 4)
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, 4, cache.Get(4))

}
