package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicates(t *testing.T) {

	assert.Equal(t, 2, LinkedListLen(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	})))

	assert.Equal(t, 3, LinkedListLen(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	})))
}
