package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getIntersectionNode(t *testing.T) {
	want := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	headA := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: want,
		},
	}

	headB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: want,
			},
		},
	}

	assert.Equal(t, want, getIntersectionNode(headA, headB))
}
