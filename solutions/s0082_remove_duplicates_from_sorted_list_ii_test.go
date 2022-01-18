package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicatesII(t *testing.T) {
	l := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					3,
					&ListNode{
						4,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
		},
	}

	assert.Equal(t, 3, LinkedListLen(deleteDuplicatesII(l)))
}
