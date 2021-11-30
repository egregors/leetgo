/*
	https://leetcode.com/problems/single-number/

	Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

	You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

package solutions

// a⊕b⊕a = (a⊕a)⊕b = 0⊕b = b
func singleNumber(nums []int) int {
	a := 0
	for _, i := range nums {
		a ^= i
	}
	return a
}
