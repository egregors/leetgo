package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reorderList(t *testing.T) {
	l := &ListNode{
		1,
		&ListNode{2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}
	reorderList(l)
	assert.Equal(t, "1 -> 4 -> 2 -> 3 -> nil", l.String())
}
