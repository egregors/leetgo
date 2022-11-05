package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie208(0)
	trie.Insert("apple")
	assert.True(t, trie.Search("apple"))
	assert.False(t, trie.Search("app"))
	assert.True(t, trie.StartsWith("app"))
	trie.Insert("app")
	assert.True(t, trie.Search("app"))

}
