package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteNode237(t *testing.T) {
	l := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	deleteNode237(l.Next)
	assert.Equal(t, "4 -> 1 -> 9 -> nil", l.String())
}
