package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseKGroup(t *testing.T) {
	l := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
	assert.Equal(t, "2 -> 1 -> 4 -> 3 -> 5 -> nil", reverseKGroup(l, 2).String())
}
