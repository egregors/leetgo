/*
	https://leetcode.com/problems/palindrome-linked-list/

	Given the head of a singly linked list, return true if it is a palindrome.
*/

package solutions

// isPalindrome234 should call isPalindrome to pass LeetCode tests
func isPalindrome234(head *ListNode) bool {
	var xs []int
	for head != nil {
		xs = append(xs, head.Val)
		head = head.Next
	}
	for i := 0; i < len(xs)/2; i++ {
		if xs[i] != xs[len(xs)-1-i] {
			return false
		}
	}
	return true
}
