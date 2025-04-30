/*
	https://leetcode.com/problems/max-consecutive-ones-iii/description

	Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array
	if you can flip at most k 0's.
*/

package solutions

func longestOnes(nums []int, k int) int {
	var l, r, cur int
	for ; r < len(nums); r++ {
		if nums[r] == 0 {
			cur++
		}

		if cur > k {
			if nums[l] == 0 {
				cur--
			}
			l++
		}
	}

	return r - l
}
