package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type iter []int

func (i iter) hasNext() bool {
	return len(i) > 0
}

func (i *iter) next() int {
	if i.hasNext() {
		res := (*i)[0]
		*i = (*i)[1:]
		return res
	}
	return -1
}

func TestNewPeekingIterator(t *testing.T) {
	peekingIterator := NewPeekingIterator(&iter{1, 2, 3}) // [1,2,3]
	assert.Equal(t, 1, peekingIterator.next())            // return 1, the pointer moves to the next element [1,2,3].
	assert.Equal(t, 2, peekingIterator.peek())            // return 2, the pointer does not move [1,2,3].
	assert.Equal(t, 2, peekingIterator.next())            // return 2, the pointer moves to the next element [1,2,3]
	assert.Equal(t, 3, peekingIterator.next())            // return 3, the pointer moves to the next element [1,2,3]
	assert.False(t, peekingIterator.hasNext())
}
