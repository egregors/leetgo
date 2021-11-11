package solutions

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	fmt.Println(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}))

	fmt.Println(deleteDuplicates(&ListNode{
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
	}))

}
