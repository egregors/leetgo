package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWordDictionary(t *testing.T) {
	wd := NewWordDictionary()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	assert.Equal(t, false, wd.Search("pad"))
	assert.Equal(t, true, wd.Search("bad"))
	assert.Equal(t, true, wd.Search(".ad"))
	assert.Equal(t, true, wd.Search("b.."))
}
