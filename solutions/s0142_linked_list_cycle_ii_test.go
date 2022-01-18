package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_detectCycle(t *testing.T) {
	cycleStart := &ListNode{
		2,
		&ListNode{
			0,
			&ListNode{
				-4,
				nil,
			},
		},
	}
	cycleStart.Next.Next.Next = cycleStart

	l := &ListNode{
		3, cycleStart,
	}
	assert.Equal(t, cycleStart, detectCycle(l))
}
