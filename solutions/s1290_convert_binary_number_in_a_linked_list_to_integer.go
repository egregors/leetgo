/*
	https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer

	Given head which is a reference node to a singly-linked list. The value of each
		node in the linked list is either 0 or 1. The linked list holds the binary
		representation of a number.

	Return the decimal value of the number in the linked list.

	The most significant bit is at the head of the linked list.
*/

package solutions

import (
	"math"
)

func getDecimalValue(head *ListNode) int {
	xs := []int{}
	for head != nil {
		xs = append(xs, head.Val)
		head = head.Next
	}

	s := 0
	for i := 0; i < len(xs); i++ {
		if xs[i] == 1 {
			s += int(math.Pow(2, float64(len(xs)-1-i)))
		}
	}

	return s
}
