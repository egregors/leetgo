package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLRUCache(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	assert.Equal(t, 1, lru.Get(1))
	lru.Put(3, 3)
	assert.Equal(t, -1, lru.Get(2))
	lru.Put(4, 4)
	assert.Equal(t, -1, lru.Get(1))
	assert.Equal(t, 3, lru.Get(3))
	assert.Equal(t, 4, lru.Get(4))
}
