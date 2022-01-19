package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_swapPairs(t *testing.T) {
	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}
	assert.Equal(t, []int{2, 1, 4, 3}, swapPairs(head).ToSlice())
}
