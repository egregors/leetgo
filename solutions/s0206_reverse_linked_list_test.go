package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {
	assert.Equal(t,
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
		reverseList(&ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		}),
	)
}
