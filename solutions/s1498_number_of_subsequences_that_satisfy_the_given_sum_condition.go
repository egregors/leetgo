/*
	You are given an array of integers nums and an integer target.

	Return the number of non-empty subsequences of nums such that the sum of the
		minimum and maximum element on it is less or equal to target. Since the answer
		may be too large, return it modulo 109 + 7.
*/

package solutions

import "sort"

func numSubseq(nums []int, target int) int {
	mod := 1_000_000_007
	n := len(nums)
	sort.Ints(nums)

	power := make([]int, n)
	power[0] = 1
	for i := 1; i < n; i++ {
		power[i] = (power[i-1] * 2) % mod
	}

	ans := 0
	l, r := 0, n-1

	for l <= r {
		if nums[l]+nums[r] <= target {
			ans += power[r-l]
			ans %= mod
			l++
		} else {
			r--
		}
	}

	return ans
}
