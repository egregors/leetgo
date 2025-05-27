/*
	https://leetcode.com/problems/minimum-size-subarray-sum/

	Given an array of positive integers nums and a positive integer target, return
		the minimal
	length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which
		the sum is
	greater than or equal to target. If there is no such subarray, return 0
		instead.
*/

package solutions

func minSubArrayLen(target int, nums []int) int {
	res := 1<<31 - 1
	sum, l := 0, 0

	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			if r-l+1 < res {
				res = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}

	if res == 1<<31-1 {
		return 0
	}

	return res
}
