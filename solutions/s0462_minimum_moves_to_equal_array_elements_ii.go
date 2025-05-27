/*
	https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/

	Given an integer array nums of size n, return the minimum number of moves
		required
	to make all array elements equal.

	In one move, you can increment or decrement an element of the array by 1.

	Test cases are designed so that the answer will fit in a 32-bit integer.
*/

package solutions

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	mid, ans := nums[len(nums)/2], 0
	for _, n := range nums {
		ans += Abs(mid - n)
	}
	return ans
}
