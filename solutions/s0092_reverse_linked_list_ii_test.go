package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseBetween(t *testing.T) {
	var ll *ListNode

	ll = NewListNode([]int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{1, 4, 3, 2, 5}, reverseBetween(ll, 2, 4).ToSlice())

	ll = NewListNode([]int{5})
	assert.Equal(t, []int{5}, reverseBetween(ll, 1, 1).ToSlice())

	ll = NewListNode([]int{3, 5})
	assert.Equal(t, []int{5, 3}, reverseBetween(ll, 1, 2).ToSlice())
}
