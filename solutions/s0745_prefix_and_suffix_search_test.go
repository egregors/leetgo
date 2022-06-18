package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWordFilter(t *testing.T) {
	/**
	 * Your WordFilter object will be instantiated and called as such:
	 * obj := Constructor(words);
	 * param_1 := obj.F(prefix,suffix);
	 */
	obj := NewWordFilter([]string{"apple"})
	assert.Equal(t, 0, obj.F("a", "e"))
}
