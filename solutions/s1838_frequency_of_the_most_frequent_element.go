/*
	https://leetcode.com/problems/frequency-of-the-most-frequent-element/description/

	The frequency of an element is the number of times it occurs in an array.

	You are given an integer array nums and an integer k. In one operation, you can choose an index of nums
	and increment the element at that index by 1.

	Return the maximum possible frequency of an element after performing at most k operations.
*/

package solutions

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	var r, l, res, curr int

	for r = 0; r < len(nums); r++ {
		target := nums[r]
		curr += target
		for (r-l+1)*target-curr > k {
			curr -= nums[l]
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}
