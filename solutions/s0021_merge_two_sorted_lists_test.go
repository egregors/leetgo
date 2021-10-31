package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	var l1, l2 *ListNode

	// l1 = [], l2 = [] => []
	l1 = nil
	l2 = nil
	assert.Nil(t, mergeTwoLists(l1, l2))

	// l1 = [], l2 = [0] => [0]
	l1 = nil
	l2 = &ListNode{Val: 0}
	assert.Equal(t, mergeTwoLists(l1, l2).String(), "0 -> nil")

	// l1 = [1,2,4], l2 = [1,3,4] => [1,1,2,3,4,4]
	l1 = intSliceToList([]int{1, 2, 4})
	l2 = intSliceToList([]int{1, 3, 4})
	assert.Equal(t, mergeTwoLists(l1, l2).String(), "1 -> 1 -> 2 -> 3 -> 4 -> 4 -> nil")
}
