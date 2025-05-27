/*
	https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/

	You are given an integer array nums and an integer x. In one operation,
	you can either remove the leftmost or the rightmost element from the array
	nums and subtract its value from x. Note that this modifies the array for
		future operations.

	Return the minimum number of operations to reduce x to exactly 0 if it is
	possible, otherwise, return -1.
*/

package solutions

import "math"

func minOperations1658(nums []int, x int) int {
	curr := 0
	for _, n := range nums {
		curr += n
	}

	n := len(nums)
	mini := math.MaxInt
	left := 0

	for right := 0; right < n; right++ {
		curr -= nums[right]
		for curr < x && left <= right {
			curr += nums[left]
			left++
		}
		if curr == x {
			mini = Minimum(mini, n-1-right+left)
		}
	}
	if mini != math.MaxInt {
		return mini
	}
	return -1
}
