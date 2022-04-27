package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotateRight(t *testing.T) {
	var l, r *ListNode

	l = NewListNode([]int{1, 2, 3, 4, 5})
	r = rotateRight(l, 2)
	assert.Equal(t, "4 -> 5 -> 1 -> 2 -> 3 -> nil", r.String())

	l = NewListNode([]int{0, 1, 2})
	r = rotateRight(l, 4)
	assert.Equal(t, "2 -> 0 -> 1 -> nil", r.String())

	l = NewListNode([]int{1, 2})
	r = rotateRight(l, 1)
	assert.Equal(t, "2 -> 1 -> nil", r.String())
}
