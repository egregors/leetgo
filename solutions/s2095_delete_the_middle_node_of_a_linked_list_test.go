package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteMiddle(t *testing.T) {
	l := NewListNode([]int{1, 3, 4, 7, 1, 2, 6})
	res := deleteMiddle(l)
	assert.Equal(t, "1 -> 3 -> 4 -> 1 -> 2 -> 6 -> nil", res.String())

	l = NewListNode([]int{1, 2, 3, 4})
	res = deleteMiddle(l)
	assert.Equal(t, "1 -> 2 -> 4 -> nil", res.String())
}
